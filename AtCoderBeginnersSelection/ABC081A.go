package main

import (
	"fmt"
	"strconv"
)

func main() {
	// {s_i| 0, 1}
	var str string
	fmt.Scan(&str)

	var ret = 0
	for i := 0; i < 3; i++ {
		// "123"のような文字列を1文字ごとに分割して整数に変換
		// str[1]をそのまな表示すると例えば1は49(文字列1はASCIIコードで49）になるので変換をかませている
		s1, _ := strconv.Atoi(string(str[i]))
		ret += s1
	}

	fmt.Println(ret)
}
