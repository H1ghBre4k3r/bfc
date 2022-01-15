package parser

const (
	ADD   = iota
	MOVE  = iota
	LOOP  = iota
	PRINT = iota
	READ  = iota
)

type Instruction struct {
	Operation int
	Operand   interface{}
}
