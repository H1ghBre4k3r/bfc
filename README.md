# go-bf

> ⚠️ Care! Currently, this project only works on x64 macOS.

Brainfuck simulator and compiler (currently only for macOS) written in Go.

# Prerequisites

For Development (and simulation): 

- `go` installed on your system
- `make` installed on your system

For Compilation:

- macOS (for now - sorry Linux)
- [NASM](https://www.nasm.us) in your path
- `ld` in your path

# Usage

Clone this repo and build the current version via

```sh
$ make release
```

You can now use the compiled executable under `./bin/darwin_amd64/go-bf`. It accepts several command line options:

- `-h`: display a help menu
- `-f <filename>`: specify a file to simulate/compile
- `-c`: enter compilation mode (currently macOS only)
- `-o <path>`: specify an output directory (where all temporary and helper files are stored)

# How it works

`go-bf` first lexes the provided file and removes all invalid/unimportant tokens. Then, it parses the lexed tokens and tries to build a structure similar to regular `ASTs`. On basis of this `AST`-like structure, it then generates NASM and compiles it.
