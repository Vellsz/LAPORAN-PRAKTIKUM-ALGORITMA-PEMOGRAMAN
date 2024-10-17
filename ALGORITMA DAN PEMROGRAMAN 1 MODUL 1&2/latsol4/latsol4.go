package main

import "fmt"

func main() {
	var f float64
	var c float64
	fmt.Scan(&f)
	c = (f - 32) * 5 / 9
	fmt.Print(c)
}
