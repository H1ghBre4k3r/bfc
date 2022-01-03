package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	inputPath := flag.String("f", "", "input file for interpreter/compiler")
	compile := flag.Bool("c", false, "flag for compiling instead of interpreting")

	flag.Parse()

	if *compile {
		fmt.Println("Compilation currently not supported!")
		os.Exit(-1)
	}

	fmt.Printf("%v\n", *inputPath)
}
