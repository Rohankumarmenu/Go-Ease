package main

import "fmt"

func main() {
	// colores := map[string]string{
	// 	"red":   "#ff0000",
	// 	"green": "#00ff00",
	// }
	colores := make(map[int]string)
	colores[10] = "#ff0000"
	fmt.Println(colores)
}
