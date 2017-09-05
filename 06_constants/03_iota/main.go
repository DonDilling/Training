package main

import "fmt"

const (
	a = iota
	b = iota
	c = iota
	d
	e
)

const (
	f = iota
	g
	h
)

func main() {
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(f)
	fmt.Println(g)
	fmt.Println(h)

}
