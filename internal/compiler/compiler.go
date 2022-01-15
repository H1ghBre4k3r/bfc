package compiler

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/H1ghBre4k3r/go-bf/internal/input"
	"github.com/H1ghBre4k3r/go-bf/internal/lexer"
	"github.com/H1ghBre4k3r/go-bf/internal/parser"
)

type Compiler struct {
	path       string
	outputPath string
	program    string
}

func New(inputPath string, outputPath string) *Compiler {
	return &Compiler{
		path:       inputPath,
		outputPath: outputPath,
		program:    input.ReadFile(inputPath),
	}
}

func (c *Compiler) Start() {
	lexed := lexer.Lex(c.program, c.path)
	parsed := parser.Parse(lexed, c.path)
	compiled := c.compile(parsed)
	c.saveCode(fmt.Sprintf(base, compiled))
}

var label = 0

func (c *Compiler) compile(parsed []parser.Instruction) string {
	toReturn := ""
	for index := 0; index < len(parsed); index++ {
		i := parsed[index]

		switch i.Operation {
		case parser.MOVE:
			toReturn += fmt.Sprintf("; MOVE %v\n", i.Operand.(int))
			toReturn += fmt.Sprintf("    add	\trdi, %v\n", i.Operand.(int))

		case parser.ADD:
			toReturn += fmt.Sprintf("; ADD %v\n", i.Operand.(int))
			toReturn += fmt.Sprintf("	add 	byte[rdx+rdi], %v\n", i.Operand.(int))

		case parser.LOOP:
			jumpLabel := label
			label++
			toReturn += fmt.Sprintf("; LOOP to loop_%v\n", jumpLabel)
			toReturn += fmt.Sprintf("loop_%v_start: \n", jumpLabel)
			toReturn += "	cmp 	byte[rdx+rdi], 0\n"
			toReturn += fmt.Sprintf("	je		loop_%v_end\n", jumpLabel)

			toReturn += c.compile(i.Operand.([]parser.Instruction))

			toReturn += fmt.Sprintf("	jmp		loop_%v_start\n", jumpLabel)
			toReturn += fmt.Sprintf("loop_%v_end:\n", jumpLabel)

		case parser.PRINT:
			toReturn += "; PRINT\n"
			toReturn += "	lea 	rsi, [rdx+rdi]\n"
			toReturn += "	call 	print\n"

		case parser.READ:
			// maybe later
		}
	}
	return toReturn
}

func (c *Compiler) saveCode(code string) {
	outputPath := c.path + ".asm"
	if c.outputPath != "" {
		name := filepath.Base(c.path)
		outputPath = filepath.Join(c.outputPath, name+".asm")
	}
	if err := os.WriteFile(outputPath, []byte(code), 0644); err != nil {
		fmt.Println(err.Error())
		os.Exit(-1)
	}
}
