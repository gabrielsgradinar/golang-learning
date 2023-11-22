package main

import "fmt"

func main() {
	var a interface{} = 10
	var b interface{} = 30

	showType(a)
	showType(b)

}

func showType(t interface{}) {
	fmt.Printf("O tipo da variável é %T\n", t)
}
