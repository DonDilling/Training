package main

import "fmt"

func main() {
	name := "Simon"
	fmt.Println(&name)
	fmt.Println(name)

	changeMe(&name)

	fmt.Println(&name)
	fmt.Println(name)

}

func changeMe(z *string) {
	fmt.Println(z)
	fmt.Println(*z)
	*z = "Michelle"

}
