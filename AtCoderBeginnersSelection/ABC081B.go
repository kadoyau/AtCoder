package main

import "fmt"

func main() {
	// parse input
	var N, tmp, counter int
	fmt.Scan(&N)

	a := []int{}
	for i := 0; i < N; i++ {
		fmt.Scan(&tmp)
		a = append(a, tmp)
	}

	// main logic
	for {
		// すべて偶数なのかをチェック
		for i := 0; i < N; i++ {
			if a[i]%2 != 0 {
				// 1つでも奇数でないなら処理を中断していままでの操作回数を返す
				fmt.Println(counter)
				return
			}
		}
		// すべて偶数だった場合、2でわってカウンタをアップ
		for i := 0; i < N; i++ {
			a[i] /= 2
		}
		counter++
	}

}
