package main

import (
	"flag"

	"github.com/H1ghBre4k3r/go-bf/internal/compiler"
	"github.com/H1ghBre4k3r/go-bf/internal/interpreter"
)

func main() {
	inputPath := flag.String("f", "", "input file for interpreter/compiler")
	compile := flag.Bool("c", false, "flag for compiling instead of interpreting")
	output := flag.String("o", "", "output directory for compilation")

	flag.Parse()

	if *compile {
		comp := compiler.New(*inputPath, *output)
		comp.Start()
	} else {
		inter := interpreter.New(*inputPath)
		inter.Start()
	}
}
