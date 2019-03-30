package main

import (
	"fmt"
	"strings"
)

func main() {
	// parse
	var S string
	fmt.Scan(&S)

	const DREAM = "dream"
	const DREAMER = "dreamer"
	const ERASE = "erase"
	const ERASER = "eraser"

	for {
		// ERASEとERASERはマッチしたら消せる
		if strings.HasPrefix(S, ERASER) {
			// stringの前方除去
			S = S[len(ERASER):]
			if len(S) == 0 {
				fmt.Println("YES")
				return
			}
			continue
		}
		if strings.HasPrefix(S, ERASE) {
			S = S[len(ERASE):]
			if len(S) == 0 {
				fmt.Println("YES")
				return
			}
			continue
		}
		// DREAMERと判定されるのは
		// DREAM ERASER or DREAM ERASE or DREAMER
		if strings.HasPrefix(S, DREAM+ERASER) {
			S = S[len(DREAM+ERASER):]
			if len(S) == 0 {
				fmt.Println("YES")
				return
			}
			continue
		}
		if strings.HasPrefix(S, DREAM+ERASE) {
			S = S[len(DREAM+ERASE):]
			if len(S) == 0 {
				fmt.Println("YES")
				return
			}
			continue
		}

		if strings.HasPrefix(S, DREAMER) {
			S = S[len(DREAMER):]
			if len(S) == 0 {
				fmt.Println("YES")
				return
			}
			continue
		}

		// DREAMERになるパターンはここまでで除去されているのでDREAMに一致するかを見ればいい
		if strings.HasPrefix(S, DREAM) {
			S = S[len(DREAM):]
			if len(S) == 0 {
				fmt.Println("YES")
				return
			}
			continue
		}

		// 条件が見つからなかった
		fmt.Println("NO")
		return
	}
}
