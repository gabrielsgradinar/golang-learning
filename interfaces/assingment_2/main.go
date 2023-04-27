package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("No file to read")
		os.Exit(1)
	}

	opf, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Println("Error to read file:", err)
		os.Exit(1)
	}

	io.Copy(os.Stdout, opf)

}
