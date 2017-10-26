package longestCommonSubSeq

//最长公共子序列，返回最长公共子序列

//递归求解  最坏情况下时间复杂度是 O(2^n), 建议用迭代
func LCSubseq(a, b string) string {
	alen := len(a)
	blen := len(b)
	if alen < 1 || blen < 1 {
		return ""
	}
	subseq := ""
	if a[alen-1] == b[blen-1] {
		subseq = a[alen-1:] + subseq
		subseq = LCSubseq(a[:alen-1], b[:blen-1]) + subseq
	} else {
		s1 := LCSubseq(a, b[:blen-1])
		s2 := LCSubseq(a[:alen-1], b)
		if len(s1) > len(s2) {
			subseq = s1 + subseq
		} else {
			subseq = s2 + subseq
		}
	}

	return subseq
}

// 动态规划
// 求最长公共子序列 返回最长公共子序列的长度 迭代
func LCSubseq1(a, b string) string {
	m, n := len(a)+1, len(b)+1
	c := make([][]int, m)
	d := make([][]string, m)
	for i := range c {
		c[i] = make([]int, n)
		d[i] = make([]string, n)
	}

	s := ""

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if a[i-1] == b[j-1] {
				c[i][j] = c[i-1][j-1] + 1
				d[i][j] = "↖"
			} else if c[i-1][j] > c[i][j-1] {
				c[i][j] = c[i-1][j]
				d[i][j] = "↑"
			} else {
				c[i][j] = c[i][j-1]
				d[i][j] = "←"
			}
		}
	}
	//构造最长公共子序列
	for i, j := m-1, n-1; i > 0 && j > 0; {
		if d[i][j] == "↖" {
			s = a[i-1:i] + s
			i--
			j--
		} else if d[i][j] == "↑" {
			i--
		} else {
			j--
		}
	}

	return s
}
