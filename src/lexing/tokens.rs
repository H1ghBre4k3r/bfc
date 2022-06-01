#[derive(Debug, Clone)]
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

/// Get the token associated with a certain brainfuck "character".
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

#[derive(Debug, Clone)]
pub struct Position {
    pub line: u64,
    pub column: u64,
}

#[cfg(test)]
mod tests {
    use super::{from, Position, Token};

    #[test]
    fn test_from_correct_tokens() {
        let dummy_position = Position { line: 0, column: 0 };

        assert!(match from('>', dummy_position.clone()).unwrap() {
            Token::RIGHT(_) => true,
            _ => false,
        });
        assert!(match from('<', dummy_position.clone()).unwrap() {
            Token::LEFT(_) => true,
            _ => false,
        });
        assert!(match from('+', dummy_position.clone()).unwrap() {
            Token::PLUS(_) => true,
            _ => false,
        });
        assert!(match from('-', dummy_position.clone()).unwrap() {
            Token::MINUS(_) => true,
            _ => false,
        });
        assert!(match from('[', dummy_position.clone()).unwrap() {
            Token::START_LOOP(_) => true,
            _ => false,
        });
        assert!(match from(']', dummy_position.clone()).unwrap() {
            Token::END_LOOP(_) => true,
            _ => false,
        });
        assert!(match from('.', dummy_position.clone()).unwrap() {
            Token::OUT(_) => true,
            _ => false,
        });
        assert!(match from(',', dummy_position.clone()).unwrap() {
            Token::IN(_) => true,
            _ => false,
        });
    }

    #[test]
    fn test_from_non_tokens() {
        let dummy_position = Position { line: 0, column: 0 };

        assert!(from('a', dummy_position.clone()).is_none());
        assert!(from('x', dummy_position.clone()).is_none());
        assert!(from(' ', dummy_position.clone()).is_none());
        assert!(from('?', dummy_position.clone()).is_none());
        assert!(from('ä', dummy_position.clone()).is_none());
    }
}
