package main

import "fmt"

func main() {
	var bmi, TinggiBadan, BeratBadan float64
	fmt.Print("Nilai bmi: ")
	fmt.Scan(&bmi)
	fmt.Print("Nilai tinggi badan (m): ")
	fmt.Scan(&TinggiBadan)
	BeratBadan = bmi * (TinggiBadan * TinggiBadan)
	fmt.Printf("Berat Badan Anda: %.0f kg\n", BeratBadan)
}
