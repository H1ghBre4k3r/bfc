#[derive(Debug)]
pub enum Token {
    LEFT(Position),
    RIGHT(Position),
    PLUS(Position),
    MINUS(Position),
    START_LOOP(Position),
    END_LOOP(Position),
    OUT(Position),
    IN(Position),
}

pub fn from(c: char, position: Position) -> Option<Token> {
    match c {
        '<' => Some(Token::LEFT(position)),
        '>' => Some(Token::RIGHT(position)),
        '+' => Some(Token::PLUS(position)),
        '-' => Some(Token::MINUS(position)),
        '[' => Some(Token::START_LOOP(position)),
        ']' => Some(Token::END_LOOP(position)),
        '.' => Some(Token::OUT(position)),
        ',' => Some(Token::IN(position)),
        _ => None,
    }
}

#[derive(Debug)]
pub struct Position {
    pub line: u64,
    pub column: u64,
}
