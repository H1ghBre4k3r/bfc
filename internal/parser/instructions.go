package parser

const (
	ADD        = iota
	MOVE       = iota
	START_LOOP = iota
	END_LOOP   = iota
	PRINT      = iota
	READ       = iota
)

type Instruction struct {
	Operation int
	Operand   int
}
