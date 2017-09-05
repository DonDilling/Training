package main

import "fmt"

const metersToFeet float64 = 3.2808399

func main() {
	var meters float64
	fmt.Print("Enter meter: ")
	fmt.Scan(&meters)
	feet := meters * metersToFeet
	fmt.Println(meters, "meters is", feet, "feet")
}
