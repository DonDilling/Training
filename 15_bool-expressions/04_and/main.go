package main

import "fmt"

func main() {
	x := 1
	if true && false {
		fmt.Println("this did not ran")
	}

	if true && x == 1 {
		fmt.Println("This ran")
	}
}
