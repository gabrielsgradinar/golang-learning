package main

const a = "Hello, World!"

var ( // declaração de escopo global
	b bool    = true
	c int     = 10
	d string  = "Gabriel"
	e float64 = 1.2
)

func main() {
	// var a string = "X" // escopo local
	a := "X" // declaração curta
	a = "B"  // atribuição
	println(a)
}
