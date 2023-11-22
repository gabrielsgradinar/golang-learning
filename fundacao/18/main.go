package main

import (
	"fmt"

	"18/matematica"

	"github.com/google/uuid"
)

func main() {

	soma := matematica.Soma(10, 20)
	fmt.Printf("Resultado: %v\n", soma)

	fmt.Println(matematica.A)
	fmt.Println(uuid.New())
}
