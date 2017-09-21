package main

import "fmt"

func main() {
	switch "Ma" {
	case "Da":
		fmt.Println("Hello Da")
	case "Ma":
		fmt.Println("Hello Ma")
	case "Ja":
		fmt.Println("Hello Ja")
	default:
		fmt.Println("Have you no friends?")

	}
}
