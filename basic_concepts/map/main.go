package main

import "fmt"

func main() {
	// var colors map[string]string
	// colors := make(map[string]string)

	colors := map[string]string{ // declaração de um map com as chaves do tipo string e valores do tipo string
		"red":   "#FF0000",
		"blue":  "#0000FF",
		"green": "#00FF00",
	}
	colors["white"] = "#FFFFFF"

	printMap(colors)

}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println("Hex code for", color, "is", hex)
	}
}
