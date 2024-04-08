package main

import "fmt"

func rechteck(seiteA int, seiteB int) int {
	var fläche int = seiteA * seiteB
	return fläche
}

func quader(grundfläche int, höhe int) int {
	var volumen int = grundfläche * höhe
	return volumen
}

func main() {
	var a int = 3
	var b int = 4
	var c int = 5
	fmt.Println("Fläche eines", a, "x", b, " Rechtecks: ", rechteck(a, b))

	fmt.Println("Volumen eines", a, "x", b, "Quaders: ", quader(rechteck(a, b), c))
}
