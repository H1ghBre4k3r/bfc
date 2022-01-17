package lexer

import (
	"fmt"
)

var lexMap = map[rune]int{
	'<': LEFT,
	'>': RIGHT,
	'+': PLUS,
	'-': MINUS,
	'[': START_LOOP,
	']': END_LOOP,
	'.': OUT,
	',': IN,
}

type Position struct {
	Line   int
	Column int
}

type LexToken struct {
	Typ      int
	Position Position
}

func Lex(code string, filePath string) []LexToken {
	fmt.Printf("[INFO] Lexing %v\n", filePath)
	lexed := make([]LexToken, 0)
	line := 0
	column := 0
	for _, c := range code {
		// filter actual symbols from "comments"
		if val, ok := lexMap[c]; ok {
			// append actual symbols to the lexed symbols
			lexed = append(lexed, LexToken{
				Typ: val,
				Position: Position{
					Line:   line,
					Column: column,
				},
			})
		}
		if c == '\n' {
			line++
			column = 0
		} else {
			column++
		}
	}
	return lexed
}
