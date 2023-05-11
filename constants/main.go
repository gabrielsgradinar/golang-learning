package main

import "fmt"

func main() {
	const secondsInHour int = 3600

	durantion := 234

	fmt.Printf("Duration in seconds: %v\n", durantion*secondsInHour)

	// x, y := 5, 0

	// fmt.Println(x / y) runtime error

	// const a = 5
	// const b = 0

	// fmt.Println(a / b) compile time error

	const n, m int = 4, 5
	const n1, m1 = 6, 7

	const (
		min1 = -500
		min2
		min3
	)

	fmt.Println(min1, min2, min3)

	const a float64 = 5.1 // typed constant
	const b = 6.7         //untyped constant

	const j = 2.2 * 5

	fmt.Printf("%T", j)

}
