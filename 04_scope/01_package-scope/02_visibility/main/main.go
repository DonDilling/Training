package main

import (
	"fmt"

	"github.com/DD_DonDilling/Training/04_scope/01_package-scope/02_visibility/visD"
)

func main() {
	fmt.Println(visD.MyName)
	visD.Printvar()
}
