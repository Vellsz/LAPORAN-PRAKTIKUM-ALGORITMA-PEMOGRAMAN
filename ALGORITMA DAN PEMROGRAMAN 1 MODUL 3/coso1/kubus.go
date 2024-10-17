package main

import "fmt"

func main() {
	var sisi float64
	fmt.Scanln(&sisi)
	vol := (sisi * sisi * sisi)
	fmt.Println(vol)
}
