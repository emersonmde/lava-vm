package main

import (
	"fmt"
	"lava-vm/pkg/class_parser"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s <classfile>\n", os.Args[0])
		os.Exit(1)
	}

	class, err := class_parser.Parse(os.Args[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading class file: %v\n", err)
		os.Exit(1)
	}

	// Display the class contents.
	class.Display()
}
