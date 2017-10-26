package sort

/*
参考：http://hxraid.iteye.com/blog/647759
桶排序的基本思想
	   假设有一组长度为N的待排关键字序列K[1....n]。首先将这个序列划分成M个的子区间(桶) 。
	   然后基于某种映射函数 ，将待排序列的关键字k映射到第i个桶中(即桶数组B的下标 i) ，
	   那么该关键字k就作为B[i]中的元素(每个桶B[i]都是一组大小为N/M的序列)。
	   接着对每个桶B[i]中的所有元素进行比较排序(可以使用快排)。
	   然后依次枚举输出B[0]....B[M]中的全部内容即是一个有序序列。

桶排序对数据的条件有特殊要求，如果上面的分数不是从100-900，
而是从0-2亿，那么分配2亿个桶显然是不可能的。
所以桶排序有其局限性，适合元素值集合并不大的情况。
*/

type ListNode struct {
	next *ListNode
	data int
}

//桶排序
//假设数据分布在  100=<score<=900 之间，每个桶内部用链表表示，
//方法：创建801(900-100)个桶。将每个考生的分数丢进f(score)=score-100的桶中
//在数据入桶的同时插入排序。然后把各个桶中的数据合并。

//分桶数要够多，否则用单链表性能很差
func BucketSort(a []int, bucket_size int) {
	buckets := make([]ListNode, bucket_size)
	for i := range buckets {
		buckets[i] = ListNode{}
	}

	for i := 0; i < len(a); i++ {
		node := &ListNode{
			data: a[i],
		}
		index := a[i] - 100
		p := &buckets[index]
		if p.data == 0 {
			buckets[index].next = node
			buckets[index].data++
		} else {
			//链表的插入排序
			for p.next != nil && p.next.data < node.data {
				p = p.next
			}
			node.next = p.next
			p.next = node
			buckets[index].data++
		}
	}

	ii := 0
	for i := 0; i < bucket_size; i++ {
		k := buckets[i].next
		for ; k != nil; k = k.next {
			a[ii] = k.data
			ii++
		}
	}
}

//桶内结构用数组，并用二分查找找到新插入元素的位置
func BucketSort2(a []int, bucket_size int) {
	buckets := make([][]int, bucket_size)
	for i := range buckets {
		buckets[i] = make([]int, 0)
	}

	for i := 0; i < len(a); i++ {
		index := a[i] - 100
		p := buckets[index]
		if len(p) == 0 {
			buckets[index] = append(buckets[index], a[i])
		} else {
			insertIndex := BinerySearch(buckets[index], a[i])
			buckets[index] = append(buckets[index], a[i])
			insertAfter(buckets[index], insertIndex+1)
			buckets[index][insertIndex] = a[i]
		}
	}

	ii := 0
	for i := 0; i < bucket_size; i++ {
		for _, v := range buckets[i] {
			a[ii] = v
			ii++
		}
	}
}

func insertAfter(a []int, index int) {
	for i := len(a) - 1; i > index; i-- {
		a[i] = a[i-1]
	}
}

//6. 给定一个有序（非降序）数组A，可含有重复元素，求最大的i使得A[i]小于等于target，不存在则返回-1
//BinerySearch Max Index Less Than or equal Target
func BinerySearch(arry []int, target int) int {
	len := len(arry)
	l := 0
	r := len - 1

	for l <= r {
		m := (l + r) / 2
		if arry[m] <= target {
			l = m + 1
		} else {
			r = m - 1
		}
	}
	if l > 0 && l < len+1 && arry[l-1] <= target {
		return l - 1
	}
	return -1

}
