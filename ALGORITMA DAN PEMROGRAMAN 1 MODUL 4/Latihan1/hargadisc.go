package main

import "fmt"

func main() {
	var TotalBelanja, Diskon int
	fmt.Print("Total belanja: ")
	fmt.Scan(&TotalBelanja)
	fmt.Print("Diskon: ")
	fmt.Scan(&Diskon)
	Diskon = TotalBelanja - (TotalBelanja * Diskon / 100)
	fmt.Print("Total harga belanja setelah diskon: ", Diskon)
}
