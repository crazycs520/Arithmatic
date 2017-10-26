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
}

func TestBinerySearchMinIEqualT(t *testing.T) {
	for _, c := range cases {
		got := BinerySearchMinIEqualT(c.arry, c.target)
		if got != c.out {
			t.Errorf("BinerySearch(%v,%d)=%d,but want %d\n", c.arry, c.target, got, c.out)
		} else {
			t.Logf("BinerySearch2(%v,%d)=%d, want %d\n", c.arry, c.target, got, c.out)
		}
	}

	for _, c := range cases1 {
		got := BinerySearchMinIEqualT(c.arry, c.target)
		if got != c.out {
			t.Errorf("BinerySearch(%v,%d)=%d,but want %d\n", c.arry, c.target, got, c.out)
		} else {
			t.Logf("BinerySearch2(%v,%d)=%d, want %d\n", c.arry, c.target, got, c.out)
		}
	}

}

func TestBinerySearchMaxIEqualT(t *testing.T) {
	for _, c := range cases {
		got := BinerySearchMaxIEqualT(c.arry, c.target)
		if got != c.out {
			t.Errorf("BinerySearch2(%v,%d)=%d,but want %d\n", c.arry, c.target, got, c.out)
		} else {
			t.Logf("BinerySearch2(%v,%d)=%d, want %d\n", c.arry, c.target, got, c.out)
		}
	}
	for _, c := range cases2 {
		got := BinerySearchMaxIEqualT(c.arry, c.target)
		if got != c.out {
			t.Errorf("BinerySearch2(%v,%d)=%d,but want %d\n", c.arry, c.target, got, c.out)
		} else {
			t.Logf("BinerySearch2(%v,%d)=%d, want %d\n", c.arry, c.target, got, c.out)
		}
	}

}

func TestBinerySearchMaxILTT(t *testing.T) {
	for _, c := range cases3 {
		got := BinerySearchMaxILTT(c.arry, c.target)
		if got != c.out {
			t.Errorf("BinerySearch(%v,%d)=%d,but want %d\n", c.arry, c.target, got, c.out)
		} else {
			t.Logf("BinerySearch2(%v,%d)=%d, want %d\n", c.arry, c.target, got, c.out)
		}
	}

}

func TestBinerySearchMinIMTT(t *testing.T) {
	for _, c := range cases4 {
		got := BinerySearchMinIMTT(c.arry, c.target)
		if got != c.out {
			t.Errorf("BinerySearch(%v,%d)=%d,but want %d\n", c.arry, c.target, got, c.out)
		} else {
			t.Logf("BinerySearch2(%v,%d)=%d, want %d\n", c.arry, c.target, got, c.out)
		}
	}

}
