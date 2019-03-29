package main

import "fmt"

func main() {
	const YUKICHI = 10000
	const ICHIYO = 5000
	const HIDEYO = 1000
	// parse input
	var N, Y int
	fmt.Scan(&N, &Y)

	// main logic
	for x := 0; x <= N; x++ {
		for y := 0; y <= N-x; y++ {
			z := N - x - y
			if x*YUKICHI+y*ICHIYO+z*HIDEYO-Y == 0 {
				fmt.Println(x, y, z)
				return
			}
		}
	}

	// 嘘だった
	fmt.Println(-1, -1, -1)

}
