package main

func main() {
	// Memória -> Endereço -> Valor
	// * indica o uso de ponteiros
	// o ponteiro é o endereçamento da memória
	a := 10
	var ponteiro *int = &a
	*ponteiro = 20
	b := &a
	*b = *b * 2

	print(a)
}
