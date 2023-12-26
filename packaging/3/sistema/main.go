package main

import (
	"fmt"

	"github.com/gabrielsgradinar/golang-learning/packaging/3/math"
)

func main(){
	m := math.NewMath(2, 3)
	fmt.Println(m.Add())
}