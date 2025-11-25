package main

import (
	"fmt"
	"os"

	"flint/internal/cli"
	"flint/internal/lexer"
	"flint/internal/parser"
	"flint/internal/typechecker"
	"flint/pkg/flint"
)

func main() {
	if len(os.Args) < 2 {
		cli.PrintHelp()
		return
	}

	switch os.Args[1] {
	case "help":
		cli.PrintHelp()
		return
	case "version":
		fmt.Print(flint.FullVersion())
		return
	case "run":
		if len(os.Args) < 3 {
			cli.Fatal("run: missing file argument")
		}
		runFile(os.Args[2])
		return
	case "compile":
		if len(os.Args) < 3 {
			cli.Fatal("run: missing file argument")
		}
		compileFile(os.Args[2])
		return
	case "check":
		if len(os.Args) < 3 {
			cli.Fatal("check: missing file argument")
		}
		checkFile(os.Args[2])
		return
	default:
		cli.Error("unknown command: " + os.Args[1])
		cli.PrintHelp()
		return
	}
}

func loadAndParse(filename string) ([]parser.Expr, *typechecker.TypeChecker) {
	src, err := os.ReadFile(filename)
	if err != nil {
		cli.Fatal(fmt.Sprintf("error reading %s: %v", filename, err))
	}
	tokens, err := lexer.Tokenize(string(src), filename)
	if err != nil {
		cli.Fatal(err.Error())
	}
	p := parser.New(tokens)
	prog, errs := p.ParseProgram()
	if len(errs) > 0 {
		for _, e := range errs {
			cli.Error(e)
		}
		os.Exit(1)
	}
	tc := typechecker.New()
	for _, ex := range prog.Exprs {
		if _, err := tc.CheckExpr(ex); err != nil {
			cli.Fatal("Type error: " + err.Error())
		}
	}
	return prog.Exprs, tc
}

func runFile(filename string) {
	fmt.Println("Running " + filename)
	_, _ = loadAndParse(filename)
	fmt.Println("Program executed (no backend yet)")
}

func compileFile(filename string) {
	fmt.Println("Compiling " + filename)
	_, _ = loadAndParse(filename)
	fmt.Println("Program compiled (no backend yet)")
}

func checkFile(filename string) {
	fmt.Println("Type checking " + filename)
	_, _ = loadAndParse(filename)
	fmt.Println("No type errors found")
}
