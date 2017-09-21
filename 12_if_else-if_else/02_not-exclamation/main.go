package main

// ! means NOT = (!true means not true) / (!false means not false) etc.

import "fmt"

func main() {
	if !true {
		fmt.Println("not ran")
	}
	if !false {
		fmt.Println("run")
	}
}
