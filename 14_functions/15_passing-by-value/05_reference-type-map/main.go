package main

import "fmt"

func main() {
	m := make(map[string]int)
	changeMe(m)
	fmt.Println(m["Simon"])
	fmt.Println(m["Michelle"])
}

func changeMe(z map[string]int) {
	z["Simon"] = 44
	z["Michelle"] = 45
}
