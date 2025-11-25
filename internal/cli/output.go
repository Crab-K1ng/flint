package cli

import (
	"fmt"
	"os"
)

func Error(msg string) {
	fmt.Fprintln(os.Stderr, msg)
}

func Fatal(msg string) {
	Error(msg)
	os.Exit(1)
}
