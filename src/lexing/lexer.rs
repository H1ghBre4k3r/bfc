use crate::lexing::tokens::{from, Position};

use super::tokens::Token;

pub fn lex(code: &String, filepath: &String) -> Vec<Token> {
    println!("[INFO] Lexing {}", filepath);
    let mut lexed: Vec<Token> = vec![];
    let mut line = 0;
    let mut column = 0;

    for c in code.chars() {
        println!("{}", c);

        if let Some(token) = from(c, Position { column, line }) {
            lexed.push(token);
        }

        if c == '\n' {
            line += 1;
            column = 0;
        } else {
            column += 1;
        }
    }

    return lexed;
}
