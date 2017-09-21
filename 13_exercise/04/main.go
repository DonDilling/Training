package main

import "fmt"

func main() {
	var x rune
	var y rune
	fmt.Print("Enter Small Number: ")
	fmt.Scan(&x)
	fmt.Print("Enter Larger number: ")
	fmt.Scan(&y)
	a := y / x
	b := y % x
	fmt.Println(y, "/", x, "=", a)
	fmt.Println("Remainder: ", b)
}
