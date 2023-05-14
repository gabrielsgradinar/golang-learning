package main

import "fmt"

func main() {
	language := "Java"

	switch language {
	case "python":
		fmt.Println("You are learning Python!")
	case "Go", "golang":
		fmt.Println("You are learning Golang!!!")
	default:
		fmt.Println("Any other programming language is a good start!")
	}
}
