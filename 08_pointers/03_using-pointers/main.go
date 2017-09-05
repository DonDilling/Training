package main

import "fmt"

func main() {
	a := 43

	fmt.Println(a)
	fmt.Println(&a)

	var b *int = &a //referencing (can be used without *int)

	fmt.Println(b)
	fmt.Println(*b) //dereferencing

	*b = 42 //changing value a; a looking to memory address b (*int) = 42
	fmt.Println(a)
}
