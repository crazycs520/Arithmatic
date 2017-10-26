package midoftwo

import "testing"

type cases struct {
	nums1, nums2 []int
	out          float64
}

var case1 = []cases{
	{[]int{1, 2, 3}, []int{4, 5, 6, 7}, 4},
	{[]int{1, 2, 3, 4}, []int{5, 6, 7}, 4},

	{[]int{}, []int{4, 5, 7, 8}, 6},
	{[]int{4, 5, 7, 8}, []int{}, 6},

	{[]int{8, 9, 10}, []int{4, 5, 6, 7}, 7},
	{[]int{5, 15, 25, 35, 45}, []int{0, 10, 20, 30, 40, 50}, 25},
}

func TestFindMedianSortedArrays(t *testing.T) {
	for _, c := range case1 {
		got := FindMedianSortedArrays(c.nums1, c.nums2)
		if got != c.out {
			t.Errorf("FindMedianSortedArrays(%v,%v)=%f,but want %f\n", c.nums1, c.nums2, got, c.out)
		} else {
			t.Logf("FindMedianSortedArrays(%v,%v)=%f, want %f\n", c.nums1, c.nums2, got, c.out)
		}

	}
}

func TestFindMedianSortedArrays2(t *testing.T) {
	for _, c := range case1 {
		got := FindMedianSortedArrays2(c.nums1, c.nums2)
		if got != c.out {
			t.Errorf("FindMedianSortedArrays2(%v,%v)=%f,but want %f\n", c.nums1, c.nums2, got, c.out)
		} else {
			t.Logf("FindMedianSortedArrays2(%v,%v)=%f, want %f\n", c.nums1, c.nums2, got, c.out)
		}

	}
}

func TestFindMedianSortedArrays3(t *testing.T) {
	for _, c := range case1 {
		got := FindMedianSortedArrays3(c.nums1, c.nums2)
		if got != c.out {
			t.Errorf("FindMedianSortedArrays3(%v,%v)=%f,but want %f\n", c.nums1, c.nums2, got, c.out)
		} else {
			t.Logf("FindMedianSortedArrays3(%v,%v)=%f, want %f\n", c.nums1, c.nums2, got, c.out)
		}

	}
}
