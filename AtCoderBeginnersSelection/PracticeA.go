package main

import (
	"fmt"
)

func main() {
	var a, b, c int
	var str string
	fmt.Scan(&a)
	fmt.Scan(&b, &c)
	fmt.Scan(&str)

	fmt.Println(a+b+c, str)
}
