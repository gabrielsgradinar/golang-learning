package main

import "fmt"

func main() {

	a1 := [4]int{10, 20, 30, 40}
	s1, s2 := a1[0:2], a1[1:3]

	fmt.Println(a1, s1, s2)

	a1[1] = 2

	fmt.Println(a1, s1, s2)
}
