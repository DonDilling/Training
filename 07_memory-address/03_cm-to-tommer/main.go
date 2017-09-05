package main

import "fmt"

const cmToInches float64 = 0.393700787

func main() {
	var cm float64
	fmt.Print("Enter cm: ")
	fmt.Scan(&cm)
	inches := cm * cmToInches
	fmt.Println(cm, "cm is", inches, "inches")

}
