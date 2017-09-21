package main

import "fmt"

func main() {
	switch "Ma" {
	case "Ta", "Ja":
		fmt.Println("Hello Ta, or, err, Ja")
	case "Me", "Ma":
		fmt.Println("Both starts with M")
	case "Fa", "Sa":
		fmt.Println("Helle Fa / Sa")
	}
}
