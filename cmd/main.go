package main

import (
	"fmt"
	"lava-vm/pkg/class"
	"lava-vm/pkg/execution_engine"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s <classfile>\n", os.Args[0])
		os.Exit(1)
	}

	class, err := class.Parse(os.Args[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading class file: %v\n", err)
		os.Exit(1)
	}

	executionEngine := execution_engine.NewExectuionEngine(class)
	err = executionEngine.Execute()

	//fmt.Printf("\n\nClass: %s\n", class)
}
