package main

import "fmt"

func main() {
	var numbers [4]int

	fmt.Printf("%v\n", numbers)
	fmt.Printf("%#v\n", numbers)

	var a1 = [4]float64{}
	fmt.Printf("%v\n", a1)

	a2 := [3]int{-10, 2, 200}
	fmt.Printf("%v\n", a2)

	//ellipsis operator
	a3 := [...]int{1, 2, 3, 5}
	fmt.Printf("a3: %v\n", a3)

	for i := range a3 {
		fmt.Printf("i: %v\n", a3[i])
	}

}
