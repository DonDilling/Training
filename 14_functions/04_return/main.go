package main

import "fmt"

func main() {
	s := greet("John ", "Doe")
	fmt.Println(s)
	fmt.Println(greet("John ", "Doe"))

}

func greet(fname, lname string) string {
	return fmt.Sprint(fname, lname)
}
