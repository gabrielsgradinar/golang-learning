package main

import "fmt"

func main() {

	salarios := map[string]int{"Gabriel": 3500, "Julia": 5700}

	delete(salarios, "Gabriel")
	salarios["Igor"] = 4000

	// fmt.Println(salarios["Igor"])

	// sal := make(map[string]int)
	// sal := map[string]int{}

	for nome, salario := range salarios {
		fmt.Printf("O salario de %s é %d\n", nome, salario)
	}

	for _, salario := range salarios {
		fmt.Printf("O salario é %d\n", salario)
	}
}
