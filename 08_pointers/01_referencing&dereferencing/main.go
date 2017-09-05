package main

import "fmt"

func main() {
	a := 43

	fmt.Println(a)
	fmt.Println(&a)

	var b *int = &a //referencing

	fmt.Println(b)
	fmt.Println(*b) //dereferencing

}
