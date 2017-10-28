package search

import (
	"testing"
)

type testCase struct {
	arry        []int
	target, out int
}

var cases = []testCase{
	{[]int{1, 2, 3, 4, 5}, 1, 0},
	{[]int{1, 2, 3, 4, 5}, 2, 1},
	{[]int{1, 2, 3, 4, 5}, 3, 2},
	{[]int{1, 2, 3, 4, 5}, 4, 3},
	{[]int{1, 2, 3, 4, 5}, 5, 4},
	{[]int{1, 2, 3, 4, 5}, 0, -1},
	{[]int{1, 2, 3, 4, 5}, 6, -1},

	{[]int{1, 2, 3, 4}, 1, 0},
	{[]int{1, 2, 3, 4}, 2, 1},
	{[]int{1, 2, 3, 4}, 3, 2},
	{[]int{1, 2, 3, 4}, 4, 3},
	{[]int{1, 2, 3, 4}, 0, -1},
	{[]int{1, 2, 3, 4}, 5, -1},

	{[]int{1, 2, 3}, 1, 0},
	{[]int{1, 2, 3}, 2, 1},
	{[]int{1, 2, 3}, 3, 2},

	{[]int{1, 2}, 1, 0},
	{[]int{1, 2}, 2, 1},

	{[]int{}, 2, -1},
}

var cases1 = []testCase{
	{[]int{1, 1, 1, 1, 1}, 1, 0},
	{[]int{2, 2, 2, 2, 2}, 2, 0},
	{[]int{3, 3, 3, 3, 3}, 3, 0},
	{[]int{4, 4, 4, 4, 4}, 4, 0},
	{[]int{5, 5, 5, 5, 5}, 5, 0},

	{[]int{1, 1, 1, 1}, 1, 0},
	{[]int{2, 2, 2, 2}, 2, 0},
	{[]int{3, 3, 3, 3}, 3, 0},
	{[]int{4, 4, 4, 4}, 4, 0},

	{[]int{1, 1, 1}, 1, 0},
	{[]int{2, 2, 2}, 2, 0},

	{[]int{1, 1}, 1, 0},
	{[]int{2, 2}, 2, 0},

	{[]int{1, 1, 1, 1, 1}, 1, 0},
	{[]int{1, 2, 2, 3, 5}, 2, 1},
	{[]int{1, 2, 3, 3, 3}, 3, 2},
	{[]int{1, 2, 3, 4, 4}, 4, 3},
}

var cases2 = []testCase{
	{[]int{1, 1, 1, 1, 1}, 1, 4},
	{[]int{2, 2, 2, 2, 2}, 2, 4},
	{[]int{3, 3, 3, 3, 3}, 3, 4},
	{[]int{4, 4, 4, 4, 4}, 4, 4},
	{[]int{5, 5, 5, 5, 5}, 5, 4},

	{[]int{1, 1, 1, 1}, 1, 3},
	{[]int{2, 2, 2, 2}, 2, 3},
	{[]int{3, 3, 3, 3}, 3, 3},
	{[]int{4, 4, 4, 4}, 4, 3},

	{[]int{1, 1, 1}, 1, 2},
	{[]int{2, 2, 2}, 2, 2},

	{[]int{1, 1}, 1, 1},
	{[]int{2, 2}, 2, 1},

	{[]int{1, 1, 1, 1, 1}, 1, 4},
	{[]int{1, 2, 2, 3, 5}, 2, 2},
	{[]int{1, 2, 3, 3, 3}, 3, 4},
	{[]int{1, 2, 3, 4, 4}, 4, 4},
}

var cases3 = []testCase{
	{[]int{1, 2, 3, 4, 5}, 1, -1},
	{[]int{1, 2, 3, 4, 5}, 2, 0},
	{[]int{1, 2, 3, 4, 5}, 3, 1},
	{[]int{1, 2, 3, 4, 5}, 4, 2},
	{[]int{1, 2, 3, 4, 5}, 5, 3},
	{[]int{1, 2, 3, 4, 5}, 6, 4},
	{[]int{1, 2, 3, 4, 5}, 0, -1},

	{[]int{1, 2, 3, 5, 5}, 4, 2},

	{[]int{1, 2, 3, 4}, 1, -1},
	{[]int{1, 2, 3, 4}, 2, 0},
	{[]int{1, 2, 3, 4}, 3, 1},
	{[]int{1, 2, 3, 4}, 4, 2},
	{[]int{1, 2, 3, 4}, 5, 3},
	{[]int{1, 2, 3, 4}, 0, -1},

	{[]int{1, 2, 4, 4}, 3, 1},

	{[]int{1, 2, 3}, 1, -1},
	{[]int{1, 2, 3}, 2, 0},
	{[]int{1, 2, 3}, 3, 1},
	{[]int{1, 2, 3}, 4, 2},

	{[]int{1, 3, 3}, 2, 0},

	{[]int{1, 2}, 1, -1},
	{[]int{1, 2}, 2, 0},
	{[]int{1, 2}, 3, 1},

	{[]int{1, 3}, 2, 0},

	{[]int{}, 2, -1},
}

var cases4 = []testCase{
	{[]int{1, 2, 3, 4, 5}, 0, 0},
	{[]int{1, 2, 3, 4, 5}, 1, 1},
	{[]int{1, 2, 3, 4, 5}, 2, 2},
	{[]int{1, 2, 3, 4, 5}, 3, 3},
	{[]int{1, 2, 3, 4, 5}, 4, 4},
	{[]int{1, 2, 3, 4, 5}, 5, -1},
	{[]int{1, 2, 3, 4, 5}, 6, -1},

	{[]int{1, 2, 3, 5, 5}, 4, 3},

	{[]int{1, 2, 3, 4}, 0, 0},
	{[]int{1, 2, 3, 4}, 1, 1},
	{[]int{1, 2, 3, 4}, 2, 2},
	{[]int{1, 2, 3, 4}, 3, 3},
	{[]int{1, 2, 3, 4}, 4, -1},
	{[]int{1, 2, 3, 4}, 5, -1},

	{[]int{1, 2, 5, 5}, 4, 2},

	{[]int{1, 2, 3}, 1, 1},
	{[]int{1, 2, 3}, 2, 2},
	{[]int{1, 2, 3}, 3, -1},
	{[]int{1, 2, 3}, 4, -1},

	{[]int{1, 3, 3}, 2, 1},

	{[]int{1, 2}, 1, 1},
	{[]int{1, 2}, 2, -1},
	{[]int{1, 2}, 3, -1},

	{[]int{1, 3}, 2, 1},

	{[]int{}, 2, -1},
}

var cases6 = []testCase{
	{[]int{1, 2, 3, 4, 5}, 1, 0},
	{[]int{1, 2, 3, 4, 5}, 2, 1},
	{[]int{1, 2, 3, 4, 5}, 3, 2},
	{[]int{1, 2, 3, 4, 5}, 4, 3},
	{[]int{1, 2, 3, 4, 5}, 5, 4},
	{[]int{1, 2, 3, 4, 5}, 6, 4},
	{[]int{1, 2, 3, 4, 5}, 0, -1},

	{[]int{1, 2, 3, 5, 5}, 4, 2},

	{[]int{1, 2, 3, 4}, 1, 0},
	{[]int{1, 2, 3, 4}, 2, 1},
	{[]int{1, 2, 3, 4}, 3, 2},
	{[]int{1, 2, 3, 4}, 4, 3},
	{[]int{1, 2, 3, 4}, 5, 3},
	{[]int{1, 2, 3, 4}, 0, -1},

	{[]int{1, 2, 4, 4}, 3, 1},

	{[]int{1, 2, 3}, 1, 0},
	{[]int{1, 2, 3}, 2, 1},
	{[]int{1, 2, 3}, 3, 2},
	{[]int{1, 2, 3}, 4, 2},

	{[]int{1, 3, 3}, 2, 0},

	{[]int{1, 2}, 1, 0},
	{[]int{1, 2}, 2, 1},
	{[]int{1, 2}, 3, 1},

	{[]int{1, 3}, 2, 0},

	{[]int{}, 2, -1},
}

//1. 给定一个有序（非降序）数组A，可含有重复元素，求最小的i使得A[i]等于target，不存在则返回-1
//BinerySearch Min Index Equal Target
func TestBinerySearchMinIEqualT(t *testing.T) {
	for _, c := range cases {
		got := BinerySearchMinIEqualT(c.arry, c.target)
		if got != c.out {
			t.Errorf("BinerySearch(%v,%d)=%d,but want %d\n", c.arry, c.target, got, c.out)
		}
	}

	for _, c := range cases1 {
		got := BinerySearchMinIEqualT(c.arry, c.target)
		if got != c.out {
			t.Errorf("BinerySearch(%v,%d)=%d,but want %d\n", c.arry, c.target, got, c.out)
		}
	}

}

//2. 给定一个有序（非降序）数组A，可含有重复元素，求最大的i使得A[i]等于target，不存在则返回-1
//BinerySearch Max Index Equal Target
func TestBinerySearchMaxIEqualT(t *testing.T) {
	for _, c := range cases {
		got := BinerySearchMaxIEqualT(c.arry, c.target)
		if got != c.out {
			t.Errorf("BinerySearch2(%v,%d)=%d,but want %d\n", c.arry, c.target, got, c.out)
		}
	}
	for _, c := range cases2 {
		got := BinerySearchMaxIEqualT(c.arry, c.target)
		if got != c.out {
			t.Errorf("BinerySearch2(%v,%d)=%d,but want %d\n", c.arry, c.target, got, c.out)
		}
	}

}

//3. 给定一个有序（非降序）数组A，可含有重复元素，求最大的i使得A[i]小于target，不存在则返回-1
//BinerySearch Max Index Less Than Target
func TestBinerySearchMaxILTT(t *testing.T) {
	for _, c := range cases3 {
		got := BinerySearchMaxILTT(c.arry, c.target)
		if got != c.out {
			t.Errorf("BinerySearch(%v,%d)=%d,but want %d\n", c.arry, c.target, got, c.out)
		}
	}

}

//5. 给定一个有序（非降序）数组A，可含有重复元素，求最小的i使得A[i]大于target，不存在则返回-1。
//也就是求大于target的最小元素的位置。
func TestBinerySearchMinIMTT(t *testing.T) {
	for _, c := range cases4 {
		got := BinerySearchMinIMTT(c.arry, c.target)
		if got != c.out {
			t.Errorf("BinerySearch(%v,%d)=%d,but want %d\n", c.arry, c.target, got, c.out)
		}
	}

}

//6. 给定一个有序（非降序）数组A，可含有重复元素，求最大的i使得A[i]小于等于target，不存在则返回-1
//BinerySearch Max Index Less Than or equal Target
func TestBinerySearchMaxILTOET(t *testing.T) {
	for _, c := range cases6 {
		got := BinerySearchMaxILTOET(c.arry, c.target)
		if got != c.out {
			t.Errorf("BinerySearch(%v,%d)=%d,but want %d\n", c.arry, c.target, got, c.out)
		}
	}

}
