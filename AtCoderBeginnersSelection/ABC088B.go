package main

import "fmt"

func main() {
	// parse input
	var N, ai int
	fmt.Scan(&N)

	var a []int
	for i := 0; i < N; i++ {
		fmt.Scan(&ai)
		a = append(a, ai)
	}

	// main logic
	var alice, bob, max int

	for {
		// アリスが手札を引く
		max, a = drawMaxNumber(a)
		alice += max
		if shouldLoopEnd(a) {
			break
		}

		// ボブが手札を引く
		max, a = drawMaxNumber(a)
		bob += max
		if shouldLoopEnd(a) {
			break
		}
	}

	// output
	fmt.Println(alice - bob)
}

func shouldLoopEnd(a []int) bool {
	// 要素がなくなったら終了
	if len(a) == 0 {
		return true
	}
	return false

}

func drawMaxNumber(a []int) (int, []int) {
	max := 0
	index := 0
	for i := 0; i < len(a); i++ {
		if a[i] > max {
			max = a[i]
			index = i
		}
	}

	// 引いたカード（最大のカード）をdropする
	a = dropIndexElement(a, index)

	return max, a
}

func dropIndexElement(a []int, index int) []int {
	a = append(a[:index], a[index+1:]...)
	return a
}

// another answer
// sortしたものを食わせて
//// aliceの得点はaの奇数番目の和
//for i := 0; i < len(a); i += 2 {
//	alice += a[i]
//}
//// bobはaの得点は偶数番目の和
//for i := 1; i < len(a); i += 2 {
//	bob += a[i]
//}
