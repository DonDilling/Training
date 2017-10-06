package main

import "fmt"

type customer struct {
	name string
	age  int
}

func main() {
	c1 := customer{"Simon", 22}
	fmt.Println(c1)
	fmt.Println(&c1.name)
	fmt.Println(&c1.age)

	changeMe(&c1)
	fmt.Println("func main post changeMe")
	fmt.Println(c1)
	fmt.Println(&c1.name)
	fmt.Println(&c1.age)
}

func changeMe(z *customer) {
	fmt.Println("func changeMe:")
	fmt.Println(z)
	fmt.Println(&z.name)
	fmt.Println("z.name = Michelle")
	z.name = "Michelle"
	fmt.Println(z)
	fmt.Println(&z.name)
}
