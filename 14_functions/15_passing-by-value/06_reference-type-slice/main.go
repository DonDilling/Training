package main

import "fmt"

// a slice is af reference type, dont need to pass the address, not referencetype
// like a int, sting, bool etc. need to pass the address, to change -

func main() {
	m := make([]string, 1, 25)
	fmt.Println(m)
	changeMe(m)
	fmt.Println(m)
}

func changeMe(z []string) {
	z[0] = "Simon"
	fmt.Println(z)
}
