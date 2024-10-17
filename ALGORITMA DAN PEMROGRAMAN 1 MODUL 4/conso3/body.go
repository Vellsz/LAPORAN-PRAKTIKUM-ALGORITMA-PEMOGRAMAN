package main

import "fmt"

func main() {
	var BeratBadan, TinggiBadan, BMI float64
	fmt.Print("Masukkan berat badan (kg): ")
	fmt.Scan(&BeratBadan)
	fmt.Print("Masukkan tinggi badan (m): ")
	fmt.Scan(&TinggiBadan)
	BMI = BeratBadan / (TinggiBadan * TinggiBadan)
	fmt.Printf("BMI: %.2f\n", BMI)
}
