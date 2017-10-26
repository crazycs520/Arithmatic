package sort

func ShellSort(a []int, n int) {
	for gap := n >> 1; gap > 0; gap >>= 1 {
		for i, j := gap, 0; i < n; i++ {
			temp := a[i]
			for j = i - gap; j >= 0 && a[j] > temp; j -= gap {
				a[j+gap] = a[j]
			}
			a[j+gap] = temp
		}
	}
}
