package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

func main() {
	var1, var2 := 'a', 'b'

	fmt.Printf("Type %T, Value %d\n", var1, var2)

	s1 := "Golang"
	fmt.Println(len(s1))

	s2 := "Maça"
	fmt.Println(len(s2))

	f := utf8.RuneCountInString(s2)
	fmt.Println(f)

	s3 := "I love Golang!"
	fmt.Println(s3[2:6])

	s4 := "昨夜最高"
	fmt.Println(s4[2:6])

	r4 := []rune(s4)
	fmt.Println(string(r4[0:3]))

	name := "Gabriel"

	fmt.Println(strings.ToTitle(name))

	fmt.Println(strings.EqualFold("Go", "go"))

}
