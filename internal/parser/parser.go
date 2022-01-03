package parser

import (
	"fmt"
	"os"

	"github.com/H1ghBre4k3r/go-bf/internal/tokens"
)

func Parse(lexed []int) []Instruction {
	parsed, index := parse(lexed, 0)
	if index != len(lexed) {
		fmt.Println("Bracket error!")
		os.Exit(-1)
	}
	return parsed
}

func parse(lexed []int, index int) ([]Instruction, int) {
	instructions := make([]Instruction, 0)
	i := 0

loop:
	for index < len(lexed) {
		l := lexed[index]
		insLen := len(instructions)
		switch l {
		case tokens.PLUS:
			if insLen == 0 || instructions[insLen-1].Operation != ADD {
				instructions = append(instructions, Instruction{
					Operation: ADD,
					Operand:   1,
				})
			} else {
				instructions[insLen-1].Operand += 1
			}

		case tokens.MINUS:
			if insLen == 0 || instructions[insLen-1].Operation != ADD {
				instructions = append(instructions, Instruction{
					Operation: ADD,
					Operand:   -1,
				})
			} else {
				instructions[insLen-1].Operand -= 1
			}

		case tokens.RIGHT:
			if insLen == 0 || instructions[insLen-1].Operation != MOVE {
				instructions = append(instructions, Instruction{
					Operation: MOVE,
					Operand:   1,
				})
			} else {
				instructions[insLen-1].Operand += 1
			}
		case tokens.LEFT:
			if insLen == 0 || instructions[insLen-1].Operation != MOVE {
				instructions = append(instructions, Instruction{
					Operation: MOVE,
					Operand:   -1,
				})
			} else {
				instructions[insLen-1].Operand -= 1
			}
		case tokens.OUT:
			instructions = append(instructions, Instruction{
				Operation: PRINT,
			})

		case tokens.IN:
			instructions = append(instructions, Instruction{
				Operation: READ,
			})

		case tokens.START_LOOP:
			instructions = append(instructions, Instruction{
				Operation: START_LOOP,
			})
			parsed, newIndex := parse(lexed, index+1)
			instructions = append(instructions, parsed...)
			index = newIndex + 1
			continue

		case tokens.END_LOOP:
			instructions = append(instructions, Instruction{
				Operation: END_LOOP,
			})
			break loop
		}

		index += 1
	}

	return instructions, index + i
}
