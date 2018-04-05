package find

//
func Find(a [][]int, rows, cols int, num int) bool {
	found := false
	if a != nil && rows > 0 && cols > 0 {
		r := 0
		l := cols - 1
		for r < rows && l >= 0 {
			if a[r][l] == num {
				found = true
				break
			} else if a[r][l] > num {
				l--
			} else {
				r++
			}
		}
	}
	return found
}
