package sort

func BubbleSort(a []int, n int) {
	for sorted := false; sorted == false; {
		sorted = true
		for i := 1; i < n; i++ {
			if a[i-1] > a[i] {
				a[i-1], a[i] = a[i], a[i-1]
				sorted = false
			}
		}
		n--
	}
}
