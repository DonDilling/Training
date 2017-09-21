package main

import "fmt"

func main() {
	if false {
		fmt.Println("1")
	} else if true {
		fmt.Println("2")
	} else {
		fmt.Println("3")
	}
}
