package main

import "fmt"

func main() {
	vals := []int{10, 5, 7, 3}
	fmt.Println(max(vals...))
	fmt.Println(min(vals...))
}

func max(vals ...int) int {
	mv := vals[0]
	for _, val := range vals {
		if val > mv {
			mv = val
		}
	}

	return mv
}

func min(vals ...int) int {
	mv := vals[0]
	for _, val := range vals {
		if val < mv {
			mv = val
		}
	}

	return mv
}
