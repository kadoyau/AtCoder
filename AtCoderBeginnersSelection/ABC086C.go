package main

import (
	"fmt"
)

func main() {

	// parse input
	var N, t, x, y int
	fmt.Scan(&N)

	ti := []int{0}
	xi := []int{0}
	yi := []int{0}

	for i := 0; i < N; i++ {
		fmt.Scan(&t, &x, &y)
		ti = append(ti, t)
		xi = append(xi, x)
		yi = append(yi, y)
	}

	for i := 0; i < N; i++ {
		// 時間内に絶対に行けない場合
		if abs(xi[i+1]-xi[i])+abs(yi[i+1]-yi[i]) > ti[i+1]-ti[i] {
			fmt.Println("No")
			return
		}
		// 最短時間から所要時間を引いた時間が偶数でないなら到達不可能
		if (abs(xi[i+1]-xi[i])+abs(yi[i+1]-yi[i])-(ti[i+1]-ti[i]))%2 != 0 {
			fmt.Println("No")
			return
		}
	}

	// そうでないならYES
	fmt.Println("Yes")
	return
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
