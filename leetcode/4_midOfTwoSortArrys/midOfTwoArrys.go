//4. Median of Two Sorted Arrays
package midoftwo

//类似于归并排序的思路
func FindMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	length1 := len(nums1)
	length2 := len(nums2)
	lengthall := length1 + length2
	if lengthall == 0 {
		return 0
	}
	mid := lengthall / 2
	i1, i2 := 0, 0
	var num, numl int
	for i := 0; i < mid+1; i++ {
		if i1 < length1 && i2 < length2 {
			if nums1[i1] <= nums2[i2] {
				numl = num
				num = nums1[i1]
				i1++
			} else {
				numl = num
				num = nums2[i2]
				i2++
			}
		} else if i1 < length1 {
			numl = num
			num = nums1[i1]
			i1++
		} else {
			numl = num
			num = nums2[i2]
			i2++
		}
	}
	if lengthall%2 == 0 {
		return float64(numl+num) / float64(2)
	}
	return float64(num)
}

//leetCode官方解决方法：https://leetcode.com/problems/median-of-two-sorted-arrays/solution/
func FindMedianSortedArrays2(nums1 []int, nums2 []int) float64 {
	m := len(nums1)
	n := len(nums2)
	if m > n {
		nums1, nums2 = nums2, nums1
		m, n = n, m
	}
	imin := 0
	imax := m
	halfLen := (m + n + 1) / 2
	for imin <= imax {
		i := (imin + imax) / 2
		j := halfLen - i
		if i < imax && nums2[j-1] > nums1[i] {
			imin = imin + 1 // i is too small
		} else if i > imin && nums1[i-1] > nums2[j] {
			imax = imax - 1 // i is too big
		} else {
			maxleft := 0
			if i == 0 {
				maxleft = nums2[j-1]
			} else if j == 0 {
				maxleft = nums1[i-1]
			} else {
				maxleft = max(nums1[i-1], nums2[j-1])
			}
			if (m+n)%2 == 1 {
				return float64(maxleft)
			}
			minright := 0
			if i == m {
				minright = nums2[j]
			} else if j == n {
				minright = nums1[i]
			} else {
				minright = min(nums1[i], nums2[j])
			}
			return float64(maxleft+minright) / float64(2.0)
		}
	}
	return 0.0
}

//调用FindKth函数，推荐~
//参考：http://blog.csdn.net/yutianzuijin/article/details/11499917
func FindMedianSortedArrays3(nums1 []int, nums2 []int) float64 {
	m := len(nums1)
	n := len(nums2)
	total := m + n
	if total%2 == 1 {
		return float64(FindKth(nums1, m, nums2, n, total/2+1))
	} else {
		return float64(FindKth(nums1, m, nums2, n, total/2)+FindKth(nums1, m, nums2, n, total/2+1)) / float64(2)
	}
}

//两个顺序数组，返回第K 大的数
func FindKth(a []int, m int, b []int, n, k int) int {
	if m > n {
		return FindKth(b, n, a, m, k)
	}
	if m == 0 {
		return b[k-1]
	}
	if k == 1 {
		return min(a[0], b[0])
	}
	pa := min(k/2, m)
	pb := k - pa
	if a[pa-1] < b[pb-1] {
		return FindKth(a[pa:], m-pa, b, n, k-pa)
	} else if a[pa-1] > b[pb-1] {
		return FindKth(a, m, b[pb:], n-pb, k-pb)
	} else {
		return a[pa-1]
	}
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
