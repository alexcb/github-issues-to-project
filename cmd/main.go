package main

import (
	"fmt"
	"os"
)

func die(msg string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, msg, args...)
	os.Exit(1)
}

func main() {
	progName := "gh-issue-projector"
	if len(os.Args) > 0 {
		progName = os.Args[0]
	}

	fmt.Printf("hello world from %s\n", progName)
}
