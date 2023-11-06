package main

import "fmt"

// array
// - tem  tamanho fixo

func main() {
	var meuArray [3]int // array de int com 3 posições
	meuArray[0] = 10
	meuArray[1] = 20
	meuArray[2] = 30
	for i, v := range meuArray {
		fmt.Printf("O valor do indice %d é %d\n", i, v)
	}

}
