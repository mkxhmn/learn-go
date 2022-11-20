package main

import (
	"fmt"
	"io"
	"os"
)

func printFileContent() {
	if len(os.Args) <= 1 {
		fmt.Println("No file name provided")
		return
	}

	fileName := os.Args[1]

	// readFile approach
	file, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println("[[ readFile approach ]]")
	fmt.Printf("%s contents:\n%s", fileName, string(file))

	// io.Copy approach
	f, errOpen := os.Open(fileName)
	if errOpen != nil {
		fmt.Println(errOpen)
		os.Exit(1)
	}

	fmt.Println("\n[[ io.Copy approach ]]")
	io.Copy(os.Stdout, f)
}
