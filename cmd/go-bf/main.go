package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/H1ghBre4k3r/go-bf/internal/interpreter"
)

func main() {
	inputPath := flag.String("f", "", "input file for interpreter/compiler")
	compile := flag.Bool("c", false, "flag for compiling instead of interpreting")

	flag.Parse()

	if *compile {
		fmt.Println("Compilation currently not supported!")
		os.Exit(-1)
	} else {
		inter := interpreter.New(*inputPath)
		inter.Start()
	}
}
