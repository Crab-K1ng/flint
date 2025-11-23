package main

import (
	"fmt"
	"os"

	"flint/internal/lexer"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("usage: flint <file>")
		return
	}

	filename := os.Args[1]

	src, err := os.ReadFile(filename)
	if err != nil {
		fmt.Printf("error reading %s: %v\n", filename, err)
		return
	}

	tokens, err := lexer.Tokenize(string(src), filename)
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, t := range tokens {
		fmt.Printf("%v(%v) @ %v:%v\n", t.Kind, t.Lexeme, t.Line, t.Column)
	}
}
