package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	var newFile *os.File
	var err error

	newFile, err = os.Create("test.txt")

	if err != nil {
		// fmt.Println(err)
		// os.Exit(1)

		log.Fatal(err)
	}

	err = os.Truncate("test.txt", 0)

	if err != nil {
		log.Fatal(err)
	}

	newFile.Close()

	file, err := os.OpenFile("test.txt", os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		log.Fatal(err)
	}
	file.Close()

	var fileInfo os.FileInfo

	fileInfo, err = os.Stat("test.txt")
	fmt.Println("File Name:", fileInfo.Name())
	fmt.Println("Size in bytes:", fileInfo.Size())
	fmt.Println("Last modified:", fileInfo.ModTime())
	fmt.Println("Is Directory ?", fileInfo.IsDir())
	fmt.Println("Permissions:", fileInfo.Mode())

	// fileInfo, err = os.Stat("eita.txt")
	// if err != nil {
	// 	if os.IsNotExist(err) {
	// 		log.Fatal("File does not exist!")
	// 	}
	// }

	oldPath := "test.txt"
	newPath := "ue.txt"
	err = os.Rename(oldPath, newPath)
	if err != nil {
		log.Fatal(err)
	}

	err = os.Remove(newPath)
	if err != nil {
		log.Fatal(err)
	}

}
