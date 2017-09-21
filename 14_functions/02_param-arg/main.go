package main

import "fmt"

func main() {
	greet("John")
	greet("Todd")
}

func greet(name string) {
	fmt.Println(name)

}
