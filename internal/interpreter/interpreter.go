package interpreter

import (
	"fmt"

	"github.com/H1ghBre4k3r/go-bf/internal/input"
	"github.com/H1ghBre4k3r/go-bf/internal/lexer"
	"github.com/H1ghBre4k3r/go-bf/internal/parser"
)

type Interpreter struct {
	path    string
	program string
}

func New(inputPath string) *Interpreter {
	return &Interpreter{
		path:    inputPath,
		program: input.ReadFile(inputPath),
	}
}

func (i *Interpreter) Start() {
	lexed := lexer.Lex(i.program, i.path)
	parsed := parser.Parse(i.path, lexed)
	i.interpret(parsed)
}

func (i *Interpreter) interpret(parsed []parser.Instruction) {
	fmt.Printf("[INFO] Interpreting %v\n", i.path)
	memory := make([]byte, 300000)
	pointer := 0
	index := 0
	eval(parsed, &index, &memory, &pointer)
}

func eval(parsed []parser.Instruction, index *int, memory *[]byte, pointer *int) {
	for ; *index < len(parsed); *index++ {
		// get current symbol
		i := parsed[*index]

		switch i.Operation {
		case parser.MOVE:
			*pointer += i.Operand.(int)

		case parser.ADD:
			(*memory)[*pointer] += byte(i.Operand.(int))

		case parser.LOOP:
			instructions := i.Operand.([]parser.Instruction)
			for (*memory)[*pointer] != 0 {
				newIndex := 0
				eval(instructions, &newIndex, memory, pointer)
			}

		case parser.PRINT:
			fmt.Print(string((*memory)[*pointer]))

		case parser.READ:
			// maybe implement that later
		}
	}
}
