package main

import "fmt"

func main() {
	// parse input
	var n, a, b int
	fmt.Scan(&n, &a, &b)

	sum := 0
	// main logic
	for i := 1; i <= n; i++ {
		digitSum := 0
		tmp := i

		for {
			//最後の桁にたどり着いたらおしまい
			if tmp == 0 {
				break
			}
			// 下一桁だけ取り出す
			digitSum += tmp % 10
			//次の桁へ
			tmp /= 10
		}

		// 各桁の我がa以上b以下のならたす
		if a <= digitSum && digitSum <= b {
			sum += i
		}
	}

	// output results
	fmt.Println(sum)
}
