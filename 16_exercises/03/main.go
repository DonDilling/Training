package main

import "fmt"

func great(numbs ...int) int {
	var greatest int
	for _, v := range numbs {
		if v > greatest {
			greatest = v
		}
	}
	return greatest
}

func main() {
	large := great(1, 5, 6, 8, 3, 4, 1, 2, 2)
	fmt.Println(large)
}
