package interpreter

import (
	"fmt"

	"github.com/H1ghBre4k3r/go-bf/internal/input"
	"github.com/H1ghBre4k3r/go-bf/internal/lexer"
	"github.com/H1ghBre4k3r/go-bf/internal/parser"
)

type Interpreter struct {
	program string
}

func New(inputPath string) *Interpreter {
	return &Interpreter{
		program: input.ReadFile(inputPath),
	}
}

func (i *Interpreter) Start() {
	lexed := lexer.Lex(i.program)
	parsed := parser.Parse(lexed)
	interpret(parsed)
}

func interpret(parsed []parser.Instruction) {
	memory := make([]byte, 300000)
	pointer := 0
	index := 0
	eval(parsed, &index, &memory, &pointer)
}

func eval(parsed []parser.Instruction, index *int, memory *[]byte, pointer *int) {
	for *index < len(parsed) {
		i := parsed[*index]
		*index++

		switch i.Operation {
		case parser.MOVE:
			*pointer += i.Operand

		case parser.ADD:
			(*memory)[*pointer] += byte(i.Operand)

		case parser.START_LOOP:
			newIndex := *index
			for (*memory)[*pointer] != 0 {
				newIndex = *index
				eval(parsed, &newIndex, memory, pointer)
			}
			*index = newIndex

		case parser.END_LOOP:
			return

		case parser.PRINT:
			fmt.Print(string((*memory)[*pointer]))
		}
	}
}
