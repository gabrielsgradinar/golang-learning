package main

import "fmt"

// Thread 1
func main(){
	canal := make(chan string) // vazio


	// Thread 2
	go func(){
		canal <- "Olá Mundo!" // cheio
	}()


	// Thread 1
	msg := <-canal // vazio
	fmt.Println(msg)
}