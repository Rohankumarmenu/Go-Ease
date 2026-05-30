package main

import "fmt"

func main() {
	colores := map[string]string{
		"red":   "#ff0000",
		"green": "#00ff00",
		"blue":  "#0000ff",
	}
	// colores := make(map[int]string)
	// colores[10] = "#ff0000"
	// delete(colores, 10)
	//fmt.Println(colores)
	printMap(colores)

}

func printMap(m map[string]string) {
	for col, hex := range m {
		fmt.Println("Hex code for", col, "is", hex)
	}
}
