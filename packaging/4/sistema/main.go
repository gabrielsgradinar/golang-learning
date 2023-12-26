package main

import (
	"fmt"

	"github.com/gabrielsgradinar/golang-learning/packaging/3/math"
	"github.com/google/uuid"
)

func main(){
	m := math.NewMath(2, 3)
	fmt.Println(m.Add())
	fmt.Println(uuid.New().String())
}