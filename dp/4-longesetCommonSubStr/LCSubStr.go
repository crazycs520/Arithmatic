package dp

// 动态规划
// 求最长公共子序列 返回最长公共子序列的长度 迭代
func LCSubStr(a, b string) (string, int) {
	m, n := len(a)+1, len(b)+1
	c := make([][]int, m)
	for i := range c {
		c[i] = make([]int, n)
	}

	biggest := 0
	lastIndex := 0

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if a[i-1] == b[j-1] {
				c[i][j] = c[i-1][j-1] + 1
				if c[i][j] > biggest {
					biggest = c[i][j]
					lastIndex = i - 1
				}
			} else {
				c[i][j] = 0
			}
		}
	}
	return a[lastIndex+1-biggest : lastIndex+1], biggest
}
