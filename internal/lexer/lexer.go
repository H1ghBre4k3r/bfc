package lexer

import "github.com/H1ghBre4k3r/go-bf/internal/tokens"

var lexMap = map[rune]int{
	'<': tokens.LEFT,
	'>': tokens.RIGHT,
	'+': tokens.PLUS,
	'-': tokens.MINUS,
	'[': tokens.START_LOOP,
	']': tokens.END_LOOP,
	'.': tokens.OUT,
	',': tokens.IN,
}

func Lex(code string) []int {
	lexed := make([]int, 0)
	for _, c := range code {
		if val, ok := lexMap[c]; ok {
			lexed = append(lexed, val)
		}
	}
	return lexed
}
