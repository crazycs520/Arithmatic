package sort

func InsertSort(a []int, n int) {
	var tmp int
	for i, j := 1, 0; i < n; i++ {
		tmp = a[i]
		for j = i - 1; j >= 0 && a[j] > tmp; j-- {
			a[j+1] = a[j]
		}
		a[j+1] = tmp
	}
}
