package sort

//当输入数据量大但数据范围小，比如[100,900]
func ArraySort(a []int, n int) {
	arry := make([]int, 901)

	for _, v := range a {
		arry[v]++
	}

	ii := 0
	for i, v := range arry {
		for j := 0; j < v; j++ {
			a[ii] = i
			ii++
		}

	}
}
