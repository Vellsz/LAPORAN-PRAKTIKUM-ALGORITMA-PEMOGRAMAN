package main

import "fmt"

func main() {
	var celsius float64
	fmt.Print("Masukan suhu dalam celcius: ")
	fmt.Scan(&celsius)
	reamur := (4 / 5) * celsius
	fahrenheit := (9/5)*celsius + 32
	kelvin := celsius + 273.15
	fmt.Print("Suhu dalam reamur: ", reamur, "\n")
	fmt.Print("Suhu dalam fahrenheit: ", fahrenheit)
	fmt.Print("Suhu dalam kelvin: ", kelvin)
}
