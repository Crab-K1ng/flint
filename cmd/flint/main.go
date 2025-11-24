package main

import (
	"fmt"
	"os"

	"flint/internal/lexer"
	"flint/internal/parser"
	"flint/internal/typechecker"
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
	p := parser.New(tokens)
	prog, errs := p.ParseProgram()
	if len(errs) > 0 {
		for _, e := range errs {
			fmt.Println(e)
		}
		return
	}
	tc := typechecker.New()
	types := make([]*typechecker.Type, len(prog.Exprs))
	for i, ex := range prog.Exprs {
		ty, err := tc.CheckExpr(ex)
		if err != nil {
			fmt.Printf("Type error: %v\n", err)
			return
		}
		types[i] = ty
	}
	for _, ex := range prog.Exprs {
		fmt.Println(parser.DumpExpr(ex))
	}
}
