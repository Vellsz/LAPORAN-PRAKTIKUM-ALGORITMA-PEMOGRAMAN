package main

import "fmt"

func main() {
	var j, v, l float64
	fmt.Print("Masukan jari-jari bola: ")
	fmt.Scan(&j)
	v = (4.0 / 3.0) * 3.1415926535 * (j * j * j)
	l = 4 * 3.1415926535 * (j * j)
	fmt.Print("Bola dengan jari-jari ", j, " cm mempunyai volume ", v, " cm^3 dan luas kulit ", l)
}
