//最长递增子序列

package lisubseq

/*
参考：http://blog.csdn.net/u013074465/article/details/45442067

1. 转化为求最长公共子序列, 将输入数组复制到另外数组后求最长公共子序列
2. 动态规划


*/

//2. 动态规划
func LISubseq(a []int) int {
	alen := len(a)
	list := make([]int, alen)
	for i := range list {
		list[i] = 1
	}

	for i := 1; i < alen; i++ {
		for j := 0; j < i; j++ {
			if a[i] > a[j] && list[i] < (list[j]+1) {
				list[i] = list[j] + 1
			}
		}
	}

	maxlist := 0
	for _, v := range list {
		if maxlist < v {
			maxlist = v
		}
	}

	return maxlist
}

//
