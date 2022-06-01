use std::{fmt::format, process};

use crate::{
    lexing::lexer::lex,
    parsing::{parser::parse, tokens::Token},
    util::read_file,
};

pub fn compile(filepath: &String) {
    let code = read_file(filepath);
    match code {
        Err(_) => {
            eprintln!("Could not open file '{}'!", filepath);
            process::exit(-1);
        }
        Ok(code) => {
            let lexed = lex(&code, filepath);
            let parsed = parse(&lexed, filepath);
            let compiled = _compile(parsed);
            println!(
                "
default rel
global _main

section .text

print:
    push    rax
    push    rdx
    push    rdi
    mov     rdi, 1 ; stdout
    mov     rax, 0x2000004 ; write
    mov     rdx, 1
    syscall
    pop     rdi
    pop     rdx
    pop     rax
    ret

_main: 
    mov     rdi, 0
    lea     rdx, array
{}
    mov     rax, 0x2000001 ; exit
    mov     rdi, 0 
    syscall

segment .bss
arraySize equ 3000
array: resb arraySize
            ",
                compiled
            )
        }
    }
}

fn _compile(parsed: Vec<Token>) -> String {
    let mut label = 1;
    let mut asm = String::from("");

    for instruction in parsed {
        match instruction {
            Token::MOVE(amount) => {
                asm += format!("; MOVE {}\n", amount).as_str();
                asm += format!("    add     \trdi, {}\n", amount).as_str();
            }
            Token::ADD(amount) => {
                asm += format!("; ADD {}\n", amount).as_str();
                asm += format!("    add     byte[rdx+rdi], {}\n", amount).as_str();
            }
            Token::LOOP(loop_content) => {
                let jump_label = label;
                label += 1;
                asm += format!("; LOOP for loop_{}\n", jump_label).as_str();
                asm += format!("loop_{}_start: \n", jump_label).as_str();
                asm += format!("    cmp     byte[rdx+rdi], 0\n").as_str();
                asm += format!("    je      loop_{}_end\n", jump_label).as_str();

                asm += _compile(loop_content).as_str();

                asm += format!("    jmp     loop_{}_start\n", jump_label).as_str();
                asm += format!("loop_{}_end: \n", jump_label).as_str();
            }
            Token::PRINT => {
                asm += format!("; PRINT\n").as_str();
                asm += format!("    lea     rsi, [rdx+rdi]\n").as_str();
                asm += format!("    call    print\n").as_str();
            }
            Token::READ => (),
        }
    }
    return asm;
}
