package compiler

const base = `
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
%v
    mov     rax, 0x2000001 ; exit
    mov     rdi, 0 
    syscall

segment .bss
arraySize equ 30000
array: resb arraySize
`
