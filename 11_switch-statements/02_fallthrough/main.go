package main

import "fmt"

func main() {
	switch "Ma" {
	case "Da":
		fmt.Println("Hello Da")
	case "Ma":
		fmt.Println("Hello Ma")
		fallthrough
	case "Ja":
		fmt.Println("Hello Ja")
		fallthrough
	case "Pa":
		fmt.Println("Hello Pa")
	default:
		fmt.Println("Have you no friends?")

	}
}
