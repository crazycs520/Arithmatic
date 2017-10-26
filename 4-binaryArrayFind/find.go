package find

func BinaryArrayFind(a [][]int, rows, cols, n int) bool {

	if a != nil && rows > 0 && cols > 0 {
		row, col := 0, cols-1

		for row < rows && col >= 0 {
			if a[row][col] == n {
				return true
			} else if a[row][col] > n {
				col--
			} else {
				row++
			}
		}
	}
	return false
}
