package main

import "fmt"

// alt+i=| so || is alt+i*2

func main() {
	if true || false {
		fmt.Println("This ran")
	}
}
