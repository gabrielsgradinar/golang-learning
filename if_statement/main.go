package main

import (
	"fmt"
	"strconv"
)

func main() {
	i, err := strconv.Atoi("45a")

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(i)
	}

	if i, err := strconv.Atoi("20"); err == nil {
		fmt.Println("No error, i is:", i)
	} else {
		fmt.Println(err)
	}
}
