package main

//Func assingt to variabel - call func-expression
//only with anonymus func - (func(){})
//only way to have func inside func

import "fmt"

func main() {
	greeting := func() {
		fmt.Println("Hello, world!")
	}
	greeting()
}
