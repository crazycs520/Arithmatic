package main

import (
	"fmt"
)

var coins = [...]int{1, 3, 5}

//动态规划
func LeastCoin(l int) int {
	min := make([]int, l+1)
	for i := range min {
		min[i] = l
	}
	min[0] = 0
	for i := 1; i <= l; i++ {
		for _, v := range coins {
			if v <= i && min[i-v]+1 < min[i] {
				min[i] = min[i-v] + 1
			}
		}
	}
	return min[l]
}

func LeastCoin2(l int) int {
	i, s := len(coins)-1, 0
	for l > 0 && i >= 0 {
		s += l / coins[i]
		l = l % coins[i]
		i--
	}

	return s
}

func main() {
	for i := 0; i < 28; i++ {
		fmt.Printf("%2d   %2d    %2d   %v\n", i, LeastCoin(i), LeastCoin2(i), LeastCoin(i) == LeastCoin2(i))
	}

}
