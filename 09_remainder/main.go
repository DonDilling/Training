package main

import "fmt"

func main() {
	x1 := 13 / 3 // division 12/3=4 (still remainder)
	x2 := 13 % 3 // tells remainder 1
	fmt.Println("result:", x1)
	fmt.Println("remainder:", x2)
	if x1 == 1 {
		fmt.Println("Odd")
	} else {
		fmt.Println("Even")
	}

	for i := 1; i < 10; i++ {
		if i%2 == 1 {
			fmt.Println("Odd")
		} else {
			fmt.Println("Even")
		}

	}
}
