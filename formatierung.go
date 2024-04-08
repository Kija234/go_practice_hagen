package main

import "fmt"

func fNachC(tempF float64) float64 {
	return (tempF - 32) / 1.8
}

func main() {
	var tempF0 float64 = -459.67
	var tempF1 float64 = 0
	var tempF2 float64 = 32
	var tempF3 float64 = 96

	fmt.Printf("%v °F entsprechen %v in °C. \n", tempF0, fNachC(tempF0))       //Formatierung durch fmt.Printf, %v findet nur hier verwendung
	fmt.Printf("%f °F entsprechen %f in °C.\n", tempF1, fNachC(tempF1))        //%f formatiert Floating direkt mit sechs Nachkommastellen
	fmt.Printf("%7.2f °F entsprechen %7.2f in °C. \n", tempF2, fNachC(tempF2)) // legt die Breite (7) und die Kommastellenfest (.2) fest

	fmt.Println(tempF1, "°F entsprechen", fNachC(tempF1), " °C.")
	fmt.Println(tempF2, "°F entsprechen", fNachC(tempF2), " °C.")
	fmt.Println(tempF3, "°F entsprechen", fNachC(tempF3), " °C.")

	//weitere Formatausgaben:
	fmt.Printf("Der Umsatz steigt um %v %% an.\n", 12.5)
}
