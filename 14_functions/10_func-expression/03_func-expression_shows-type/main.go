package main

import "fmt"

func main() {
	greeting := func() {
		fmt.Println("Hello, world!")
	}
	greeting() //Calls the variabel, which is aof type func() and runs it - Hello, world!
	fmt.Printf("%T \n", greeting)
}
