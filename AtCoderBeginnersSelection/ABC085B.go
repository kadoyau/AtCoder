package main

import "fmt"

func main() {
	// parse input
	var N, di int
	fmt.Scan(&N)

	var d []int
	for i := 0; i < N; i++ {
		fmt.Scan(&di)
		d = append(d, di)
	}

	// main logic
	uniqueD := unique(d)

	fmt.Println(len(uniqueD))

}

func unique(intSlice []int) []int {
	// 数値を格納する変数を作る
	// 格納された数値のflagはtrueにする
	keys := make(map[int]bool)
	var list []int
	for _, entry := range intSlice {
		// 格納されていなければ数値を格納してフラグをtrueに
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}
