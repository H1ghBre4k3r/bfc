package interpreter

import (
	"fmt"

	"github.com/H1ghBre4k3r/go-bf/internal/input"
	"github.com/H1ghBre4k3r/go-bf/internal/lexer"
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
	fmt.Printf("lexed: %v\n", lexed)
}
