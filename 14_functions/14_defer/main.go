package main

import "fmt"

func hello() {
	fmt.Print("Hello ")
}

func world() {
	fmt.Println("world")
}

func main() {
	defer world() //defer means - line runs just before func exits
	hello()
}
