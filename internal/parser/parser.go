package parser

import (
	"fmt"
	"reflect"

	"github.com/H1ghBre4k3r/go-bf/internal/lexer"
)

func Parse(lexed []lexer.LexToken, filePath string) []Instruction {
	fmt.Printf("[INFO] Parsing %v\n", filePath)
	parsed, _, err := parse(lexed, 0, filePath, false)
	// if we didn't parse everything before returning, there was a bracket error
	if err != nil {
		fmt.Printf("[ERROR] %v\n", err.Error())
	}
	return parsed
}

func parse(lexed []lexer.LexToken, index int, filePath string, inLoop bool) ([]Instruction, int, error) {
	instructions := make([]Instruction, 0)
	i := 0

	for index < len(lexed) {
		l := lexed[index]
		instructionCount := len(instructions)
		// switch through current
		switch l.Typ {
		case lexer.PLUS:
			if instructionCount == 0 || instructions[instructionCount-1].Operation != ADD {
				instructions = append(instructions, Instruction{
					Operation: ADD,
					Operand:   1,
				})
			} else {
				curVal, ok := instructions[instructionCount-1].Operand.(int)
				if !ok {
					panic(fmt.Sprintf("error during parsing. token of type '%v' had non working operand of type %v!", l.Typ, reflect.TypeOf(instructions[instructionCount-1].Operand).Kind()))
				}
				instructions[instructionCount-1].Operand = curVal + 1
			}

		case lexer.MINUS:
			if instructionCount == 0 || instructions[instructionCount-1].Operation != ADD {
				instructions = append(instructions, Instruction{
					Operation: ADD,
					Operand:   -1,
				})
			} else {
				curVal, ok := instructions[instructionCount-1].Operand.(int)
				if !ok {
					panic(fmt.Sprintf("error during parsing. token of type '%v' had non working operand of type %v!", l.Typ, reflect.TypeOf(instructions[instructionCount-1].Operand).Kind()))
				}
				instructions[instructionCount-1].Operand = curVal - 1
			}

		case lexer.RIGHT:
			if instructionCount == 0 || instructions[instructionCount-1].Operation != MOVE {
				instructions = append(instructions, Instruction{
					Operation: MOVE,
					Operand:   1,
				})
			} else {
				curVal, ok := instructions[instructionCount-1].Operand.(int)
				if !ok {
					panic(fmt.Sprintf("error during parsing. token of type '%v' had non working operand of type %v!", l.Typ, reflect.TypeOf(instructions[instructionCount-1].Operand).Kind()))
				}
				instructions[instructionCount-1].Operand = curVal + 1
			}

		case lexer.LEFT:
			if instructionCount == 0 || instructions[instructionCount-1].Operation != MOVE {
				instructions = append(instructions, Instruction{
					Operation: MOVE,
					Operand:   -1,
				})
			} else {
				curVal, ok := instructions[instructionCount-1].Operand.(int)
				if !ok {
					panic(fmt.Sprintf("error during parsing. token of type '%v' had non working operand of type %v!", l.Typ, reflect.TypeOf(instructions[instructionCount-1].Operand).Kind()))
				}
				instructions[instructionCount-1].Operand = curVal - 1
			}

		case lexer.OUT:
			instructions = append(instructions, Instruction{
				Operation: PRINT,
			})

		case lexer.IN:
			instructions = append(instructions, Instruction{
				Operation: READ,
			})

		case lexer.START_LOOP:
			parsed, newIndex, err := parse(lexed, index+1, filePath, true)
			if err != nil {
				return instructions, newIndex, fmt.Errorf("opening bracket not closed: \n\t%v:%v:%v", filePath, l.Position.Line, l.Position.Column)
			}
			instructions = append(instructions, Instruction{
				Operation: LOOP,
				Operand:   parsed,
			})
			index = newIndex + 1
			continue

		case lexer.END_LOOP:
			var err error
			if !inLoop {
				err = fmt.Errorf("unexpected closing bracket at: \n\t%v:%v:%v", filePath, l.Position.Line, l.Position.Column)
			}
			return instructions, index + i, err
		}

		index += 1
	}

	// check, if we are still in a loop but reached end of tokens
	var err error
	if inLoop {
		err = fmt.Errorf("expected closing bracket at: \n\t%v:%v:%v", filePath, lexed[index+i-1].Position.Line, lexed[index+i-1].Position.Column)
	}
	return instructions, index + i, err
}
