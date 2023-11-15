package main

type Number interface {
	~int | ~float64
}

func Soma[T Number](m map[string]T) T {
	var soma T
	for _, v := range m {
		soma += v
	}
	return soma
}

func Compara[T comparable](a T, b T) bool {
	if a == b {
		return true
	}
	return false
}

func main() {
	a := map[string]int{"Gabriel": 1000, "Julia": 2000, "Tiago": 3000}
	b := map[string]float64{"Gabriel": 1250.6, "Julia": 2000.50, "Tiago": 3100.20}

	println(Soma(a))
	println(Soma(b))

	println(Compara(10, 10.0))
}
