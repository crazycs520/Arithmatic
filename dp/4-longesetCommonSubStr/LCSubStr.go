package dp

// 动态规划
// 求最长公共子序列 返回最长公共子序列的长度 迭代
func LCSubStr(s1, s2 string) string {
	if len(s1) == 0 || len(s2) == 0 {
		return ""
	}
	m := make([][]int, len(s1)+1)
	for i := range m {
		m[i] = make([]int, len(s2)+1)
	}
	largest := 0
	lastIndex := 0
	for i := 1; i <= len(s1); i++ {
		for j := 1; j <= len(s2); j++ {
			if s1[i-1] == s2[j-1] {
				m[i][j] = m[i-1][j-1] + 1
				if m[i][j] > largest {
					largest = m[i][j]
					lastIndex = i - 1
				}
			} else {
				m[i][j] = 0
			}
		}
	}
	return s1[lastIndex-largest+1 : lastIndex+1]
}


func LCSubStr2(s1, s2 string) string {
	if len(s1) == 0 || len(s2) == 0 {
		return ""
	}
	start, end := 0, 0
	for i := 0; i < len(s1); i++ {
		if (len(s1) - i) < (start - end) {
			break
		}
		for i1 := 0; i1 < len(s2); i1++ {
			if (len(s2)-i1) < (end-start) {
				break
			}
			if s1[i] == s2[i1] {
				k, j := i, i1
				for k < len(s1) && j < len(s2) && s1[k] == s2[j] {
					k++
					j++
				}
				if (k - i) > (end - start) {
					start = i
					end = k
				}
			}
		}
	}
	return s1[start:end]
}

