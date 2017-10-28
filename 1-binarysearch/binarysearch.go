package search

//1. 给定一个有序（非降序）数组A，可含有重复元素，求最小的i使得A[i]等于target，不存在则返回-1
//BinerySearch Min Index Equal Target
func BinerySearchMinIEqualT(arry []int, target int) int {
	alen := len(arry)
	l := 0
	r := alen
	for l < r {
		m := (l + r) / 2
		if arry[m] >= target {
			r = m
		} else {
			l = m + 1
		}
	}
	if r < alen && arry[r] == target {
		return r
	}
	return -1
}

//2. 给定一个有序（非降序）数组A，可含有重复元素，求最大的i使得A[i]等于target，不存在则返回-1
//BinerySearch Max Index Equal Target
func BinerySearchMaxIEqualT(arry []int, target int) int {
	alen := len(arry)
	l := 0
	r := alen
	for l < r {
		m := (l + r) / 2
		if arry[m] > target {
			r = m
		} else {
			l = m + 1
		}
	}
	if r > 0 && arry[r-1] == target {
		return r - 1
	}
	return -1
}

//3. 给定一个有序（非降序）数组A，可含有重复元素，求最大的i使得A[i]小于target，不存在则返回-1
//BinerySearch Max Index Less Than Target
func BinerySearchMaxILTT(arry []int, target int) int {
	alen := len(arry)
	l := 0
	r := alen
	for l < r {
		m := (l + r) / 2
		if arry[m] >= target {
			r = m
		} else {
			l = m + 1
		}
	}
	if r > 0 && arry[r-1] < target {
		return r - 1
	}
	return -1
}

//5. 给定一个有序（非降序）数组A，可含有重复元素，求最小的i使得A[i]大于target，不存在则返回-1。
//也就是求大于target的最小元素的位置。
//BinerySearch Min Index More Than Target
func BinerySearchMinIMTT(arry []int, target int) int {
	alen := len(arry)
	l := 0
	r := alen
	for l < r {
		m := (l + r) / 2
		if arry[m] > target {
			r = m
		} else {
			l = m + 1
		}
	}
	if r < alen && arry[r] > target {
		return r
	}
	return -1
}

//6. 给定一个有序（非降序）数组A，可含有重复元素，求最大的i使得A[i]小于等于target，不存在则返回-1
//BinerySearch Max Index Less Than or equal Target
func BinerySearchMaxILTOET(arry []int, target int) int {
	alen := len(arry)
	l := 0
	r := alen
	for l < r {
		m := (l + r) / 2
		if arry[m] > target {
			r = m
		} else {
			l = m + 1
		}
	}
	l--
	return l
}
