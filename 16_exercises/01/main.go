package main

import "fmt"

func main() {
	var number int
	fmt.Print("Enter interger: ")
	fmt.Scan(&number)
	h, even := half(number)
	fmt.Println(h, even)
	j, even2 := half2(number)
	fmt.Println(j, even2)

}

func half(n int) (int, bool) {
	return n / 2, n%2 == 0
}

func half2(m int) (float64, bool) {
	return float64(m) / 2, m%2 == 0
}
