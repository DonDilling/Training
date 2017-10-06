package main

// Only changing age beacause passing the address where age is storede, not in 01_int

import "fmt"

func main() {
	age := 44
	fmt.Println(&age) //0xc420012050

	changeMe(&age)

	fmt.Println(&age) //0xc420012050
	fmt.Println(age)  //24
}

func changeMe(z *int) {
	fmt.Println(z)  //0xc420012050
	fmt.Println(*z) //44
	*z = 24
	fmt.Println(z)  //0xc420012050
	fmt.Println(*z) //24
}
