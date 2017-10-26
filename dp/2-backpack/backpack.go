package main

import (
	"fmt"
)

/*
v : 体积数组 如 []int{4, 3, 5, 2, 5}
w : 价值数组 如 []int{9, 6, 1, 4, 1}
c : 背包容量 如 10
返回值：最大价值

*/
func backpack(v []int, w []int, c int) (int, []int) {
	n := len(v)
	//声明d二维数组
	d := make([][]int, n+1)
	for i := range d {
		d[i] = make([]int, c+1)
	}

	x := make([]int, n)

	for i := 0; i <= n; i++ {
		for j := 0; j <= c; j++ {
			if i > 0 {
				d[i][j] = d[i-1][j]
				if j >= v[i-1] {
					if d[i-1][j-v[i-1]]+w[i-1] > d[i][j] {
						d[i][j] = d[i-1][j-v[i-1]] + w[i-1]
					}
				}
			} else {
				d[i][j] = 0
			}

		}
	}

	j := c
	for i := n; i > 0; i-- {
		if d[i][j] > d[i-1][j] {
			x[i-1] = 1
			j = j - v[i-1]
		}
	}

	return d[n][c], x
}

//优化空间复杂度
//如果输入是v[i] w[i] 单行输入, v , w数组也可以优化
func backpack2(v []int, w []int, c int) int {
	d := make([]int, c+1)
	n := len(v)
	for i := 0; i <= n; i++ {
		for j := c; j >= 0; j-- {
			if i > 0 && j >= v[i-1] {
				if d[j-v[i-1]]+w[i-1] > d[j] {
					d[j] = d[j-v[i-1]] + w[i-1]
				}
				fmt.Printf("%d,d[%d]=%d\n", i, j, d[j])
			}
		}
	}

	return d[c]
}

func main() {
	v := []int{4, 3, 5, 2, 5}
	w := []int{9, 6, 1, 4, 1}
	c := 10
	fmt.Println(backpack(v, w, c))
	fmt.Println(backpack2(v, w, c))
}
