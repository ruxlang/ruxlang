package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/ruxlang/ruxlang/internal/grammars/ruxlang"
)

func main() {
	args := os.Args[1:]
	if len(args) < 1 {
		log.Fatal("Please provide a file to parse.\n")
		return
	}

	path, err := filepath.Abs(args[0])
	if err != nil {
		log.Fatalf("Failed to get the file path: %s\n", err.Error())
		return
	}

	listener, err := ruxlang.Parse(path)
	if err != nil {
		log.Fatalf("Failed to open the file: %s\n", err.Error())
		return
	}

	fmt.Println("Global Imports")
	for _, imp := range listener.GlobalImports {
		fmt.Println(imp)
	}
	fmt.Println("\nNamed Imports")
	for ref, imp := range listener.NamedImports {
		fmt.Printf("%s=%s\n", ref, imp)
	}
}
