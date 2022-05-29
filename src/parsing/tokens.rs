#[derive(Debug)]
pub enum Token {
    ADD(i8),
    MOVE(i64),
    LOOP(Vec<Token>),
    PRINT,
    READ,
}
