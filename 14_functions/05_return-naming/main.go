package main

import "fmt"

func main() {
	fmt.Println(greet("John ", "Doe"))
}

func greet(fname, lname string) (s string) {
	s = fmt.Sprint(fname, lname)
	return
}

// Don't do this, do 04_return
