package sort

func QuickSort(a []int) {
	n := len(a)
	if n <= 1 {
		return
	}
	mid := a[0]
	i := 1
	head, tail := 0, n-1
	for head < tail {
		if a[i] > mid {
			a[i], a[tail] = a[tail], a[i]
			tail--
		} else {
			a[i], a[head] = a[head], a[i]
			head++
			i++
		}
	}
	a[head] = mid
	QuickSort(a[:head])
	QuickSort(a[head+1:])
}
