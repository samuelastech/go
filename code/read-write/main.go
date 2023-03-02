package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	// Has the user passed an arg?
	if len(os.Args) == 0 {
		fmt.Println("You need to provide a file")
		os.Exit(1)
	}

	// Open the file
	filename := os.Args[1]
	file, error := os.Open(filename)

	if error != nil {
		fmt.Println(error)
		os.Exit(1)
	}

	io.Copy(os.Stdout, file)
}
