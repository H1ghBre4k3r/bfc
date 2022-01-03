package input

import (
	"fmt"
	"os"
)

func ReadFile(path string) string {
	data, err := os.ReadFile(path)
	if err != nil {
		fmt.Printf("Cannot open file: '%v'\n", path)
		os.Exit(-1)
	}
	return string(data)
}
