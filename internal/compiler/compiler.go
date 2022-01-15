package compiler

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

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

	fmt.Printf("[INFO] Generating NASM for %v\n", c.path)
	compiled := c.compile(parsed)
	c.saveCode(fmt.Sprintf(base, compiled))
	c.compileNasm()
	c.linkNasm()
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
	if err := os.WriteFile(c.outPath(".asm"), []byte(code), 0644); err != nil {
		fmt.Println(err.Error())
		os.Exit(-1)
	}
}

func (c *Compiler) outPath(ending string) string {
	outputPath := c.path + ending
	if c.outputPath != "" {
		name := filepath.Base(c.path)
		outputPath = filepath.Join(c.outputPath, name+ending)
	}
	return outputPath
}

func (c *Compiler) compileNasm() {
	fmt.Printf("[INFO] Compiling NASM for %v\n", c.path)
	if err := exec.Command("nasm", "-f", "macho64", c.outPath(".asm")).Run(); err != nil {
		fmt.Printf("[ERROR] could not compile NASM (%v)\n", err.Error())
	}
}

func (c *Compiler) linkNasm() {
	fmt.Printf("[INFO] Linking NASM for %v\n", c.path)
	if err := exec.Command("ld", "-macos_version_min", "10.12.0", "-L/Library/Developer/CommandLineTools/SDKs/MacOSX.sdk/usr/lib", "-lSystem", "-o", c.executablemain(), c.outPath(".o")).Run(); err != nil {
		fmt.Printf("[ERROR] could not link NASM (%v)\n", err.Error())
	}
}

func (c *Compiler) executablemain() string {
	fileName := c.outPath("")
	if pos := strings.LastIndexByte(fileName, '.'); pos != -1 {
		return fileName[:pos]
	}
	return fileName
}
