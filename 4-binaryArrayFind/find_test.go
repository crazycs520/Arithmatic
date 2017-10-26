package find

import (
	"testing"
)

var case1 = [][]int{
	{1, 2, 8, 9},
	{2, 4, 9, 12},
	{4, 7, 10, 13},
	{6, 8, 11, 15},
}

func test(a [][]int, rows, cols, n int, result bool, t *testing.T) {
	if BinaryArrayFind(a, rows, cols, n) != result {
		t.Fatalf("BinaryArrayFind %d error\n", n)
	}
}
func TestBinaryArrayFind(t *testing.T) {
	test(case1, len(case1), len(case1[0]), 16, false, t)
	test(case1, 4, 4, 0, false, t)
	test(case1, 4, 4, 5, false, t)

	test(case1, 4, 4, 1, true, t)
	test(case1, 4, 4, 9, true, t)
	test(case1, 4, 4, 6, true, t)
	test(case1, 4, 4, 15, true, t)

	test(case1, 4, 4, 7, true, t)

}
