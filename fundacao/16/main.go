package main

import "fmt"

func main() {
	var a interface{} = "Gabriel"
	println(a.(string))

	res, ok := a.(int)
	fmt.Printf("O valor de res é %v e o resultado de ok é %v", res, ok)
}
