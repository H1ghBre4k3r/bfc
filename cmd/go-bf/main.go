package main

import (
	"flag"
	"fmt"
)

func main() {
	inputPath := flag.String("f", "", "input file for interpreter/compiler")

	flag.Parse()

	fmt.Printf("%v\n", *inputPath)
}
