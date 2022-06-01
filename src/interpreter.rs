use std::process;

use crate::{
    lexing::lexer::lex,
    parsing::{parser::parse, tokens::Token},
    util::read_file,
};

pub fn interpret(filepath: &String) {
    let code = read_file(filepath);
    if let Err(_) = code {
        eprintln!("Could not open file '{}'!", filepath);
        process::exit(-1);
    } else if let Ok(code) = code {
        let lexed = lex(&code, filepath);
        let parsed = parse(&lexed, filepath);
        _interpret(parsed, filepath);
    }
}

fn _interpret(parsed: Vec<Token>, filepath: &String) {
    println!("[INFO] Interpreting {}", filepath);
    let memory: Vec<u8> = vec![0; 300000];
    eval(&parsed, 0, memory, 0);
}

fn eval(
    parsed: &Vec<Token>,
    mut index: usize,
    mut memory: Vec<u8>,
    mut pointer: usize,
) -> (Vec<u8>, usize, usize) {
    while index < parsed.len() {
        let i = &parsed[index];

        match i {
            Token::MOVE(amount) => {
                if amount.is_negative() {
                    pointer -= amount.abs() as usize
                } else {
                    pointer += *amount as usize
                }
            }
            Token::ADD(amount) => {
                if amount.is_negative() {
                    memory[pointer] -= amount.abs() as u8;
                } else {
                    memory[pointer] += *amount as u8;
                }
            }
            Token::LOOP(instructions) => {
                while memory[pointer] != 0 {
                    let new_index = 0;
                    let (new_memory, _, new_pointer) =
                        eval(instructions, new_index, memory.clone(), pointer);
                    memory = new_memory;
                    pointer = new_pointer;
                }
            }
            Token::PRINT => print!("{}", memory[pointer] as char),
            Token::READ => (),
        }
        index += 1;
    }

    return (memory, index, pointer);
}
