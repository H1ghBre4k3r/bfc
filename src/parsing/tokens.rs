#[derive(Debug)]
pub enum Token {
    ADD(i64),
    MOVE(i64),
    LOOP(Vec<Token>),
    PRINT,
    READ,
}
