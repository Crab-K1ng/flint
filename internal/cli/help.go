package cli

import "fmt"

func PrintHelp() {
	fmt.Println("Usage:")
	fmt.Println("  flint " + "<command> [options]")
	fmt.Println()

	fmt.Println("Commands:")

	printCmd("run <file>", "Execute a Flint source file.")
	printCmd("compile <file>", "Compile a Flint file to a target backend.")
	printCmd("check <file>", "Type-check only; no execution.")
	printCmd("version", "Print the Flint compiler version.")
	printCmd("help", "Show this help message.")

	fmt.Println()
	fmt.Println("Tip: Run 'flint help <command>' for more info.")
}

func printCmd(cmd, desc string) {
	fmt.Printf("  %-14s  %s\n", cmd, desc)
}
