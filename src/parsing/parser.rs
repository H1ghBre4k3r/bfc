use std::process;

use crate::lexing;

/// Parse and optimize a list of lexed tokens.
pub fn parse(lexed: &Vec<lexing::tokens::Token>, filepath: &String) -> Vec<super::tokens::Token> {
    println!("[INFO] Parsing {}", filepath);
    match _parse(lexed, 0, filepath, false) {
        Err(reason) => {
            eprintln!("[ERROR] {}", reason);
            process::exit(-1);
        }
        Ok((parsed, _)) => return parsed,
    }
}

fn _parse(
    lexed: &Vec<lexing::tokens::Token>,
    mut index: usize,
    filepath: &String,
    in_loop: bool,
) -> Result<(Vec<super::tokens::Token>, usize), String> {
    let mut parsed: Vec<super::tokens::Token> = vec![];

    while index < lexed.len() {
        let l = &lexed[index];
        let instruction_count = parsed.len();

        match l {
            lexing::tokens::Token::LEFT(_) => {
                if instruction_count == 0
                    || match parsed[instruction_count - 1] {
                        super::tokens::Token::MOVE(_) => false,
                        _ => true,
                    }
                {
                    parsed.push(super::tokens::Token::MOVE(-1))
                } else {
                    if let super::tokens::Token::MOVE(amount) = parsed[instruction_count - 1] {
                        parsed[instruction_count - 1] = super::tokens::Token::MOVE(amount - 1);
                    }
                }
            }
            lexing::tokens::Token::RIGHT(_) => {
                if instruction_count == 0
                    || match parsed[instruction_count - 1] {
                        super::tokens::Token::MOVE(_) => false,
                        _ => true,
                    }
                {
                    parsed.push(super::tokens::Token::MOVE(1))
                } else {
                    if let super::tokens::Token::MOVE(amount) = parsed[instruction_count - 1] {
                        parsed[instruction_count - 1] = super::tokens::Token::MOVE(amount + 1);
                    }
                }
            }
            lexing::tokens::Token::PLUS(_) => {
                if instruction_count == 0
                    || match parsed[instruction_count - 1] {
                        super::tokens::Token::ADD(_) => false,
                        _ => true,
                    }
                {
                    parsed.push(super::tokens::Token::ADD(1))
                } else {
                    if let super::tokens::Token::ADD(amount) = parsed[instruction_count - 1] {
                        parsed[instruction_count - 1] = super::tokens::Token::ADD(amount + 1);
                    }
                }
            }
            lexing::tokens::Token::MINUS(_) => {
                if instruction_count == 0
                    || match parsed[instruction_count - 1] {
                        super::tokens::Token::ADD(_) => false,
                        _ => true,
                    }
                {
                    parsed.push(super::tokens::Token::ADD(-1))
                } else {
                    if let super::tokens::Token::ADD(amount) = parsed[instruction_count - 1] {
                        parsed[instruction_count - 1] = super::tokens::Token::ADD(amount - 1);
                    }
                }
            }
            lexing::tokens::Token::START_LOOP(pos) => {
                match _parse(lexed, index + 1, filepath, true) {
                    Err(_) => {
                        return Err(format!(
                            "opening bracket not closed: \n\t{}:{}:{}",
                            filepath, pos.line, pos.column
                        ))
                    }
                    Ok((new_parsed, new_index)) => {
                        parsed.push(super::tokens::Token::LOOP(new_parsed));
                        index = new_index + 1;
                        continue;
                    }
                }
            }
            lexing::tokens::Token::END_LOOP(pos) => {
                if !in_loop {
                    return Err(format!(
                        "unexpected closing bracket at: \n\t{}:{}:{}",
                        filepath, pos.line, pos.column
                    ));
                }
                return Ok((parsed, index));
            }
            lexing::tokens::Token::OUT(_) => parsed.push(super::tokens::Token::PRINT),
            lexing::tokens::Token::IN(_) => parsed.push(super::tokens::Token::READ),
        }
        index += 1
    }

    if in_loop {
        return Err(format!("expected closing bracket at end of file"));
    }

    return Ok((parsed, index));
}
