package main

import "fmt"

func main() {
	// parse input
	var a, b, c, x int
	fmt.Scan(&a)
	fmt.Scan(&b)
	fmt.Scan(&c)
	fmt.Scan(&x)

	// main logic
	// 全探索する
	var counter int
	for i := 0; i <= a; i++ {
		for j := 0; j <= b; j++ {
			for k := 0; k <= c; k++ {
				if 500*i+100*j+50*k-x == 0 {
					counter++
				}
			}
		}
	}

	fmt.Print(counter)
}
