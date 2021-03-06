package main

import "fmt"

func main() {
	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#4bf7",
	}
	fmt.Println(colors)

	colors2 := make(map[int]string)
	colors2[10] = "#ffffff"
	fmt.Println(colors2)

	delete(colors2, 10)
	fmt.Println(colors2)

	printMap(colors)
}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println("Hex code for", color, "is", hex)
	}
}
