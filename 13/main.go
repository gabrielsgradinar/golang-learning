package main

func soma(a, b *int) int {
	*a = 50
	*b = 30
	return *a + *b
}

func main() {

	a := 10
	b := 20
	println(soma(&a, &b))

	println(a)
	println(b)

}
