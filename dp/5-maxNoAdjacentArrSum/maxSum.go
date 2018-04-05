package maxNJAS

//递归
func Rec_maxSum(arr []int , i int) int{
	if i == 0 {
		return arr[i]
	}
	if i == 1 {
		return max(arr[0],arr[1])
	}
	a := Rec_maxSum(arr,i-2)+arr[i]
	b := Rec_maxSum(arr,i-1)
	return max(a,b)
}

//迭代
func Dp_maxSum(arr []int) int{
	l := len(arr)
	opt := make([]int,l)
	opt[0]=arr[0]
	opt[1]=max(arr[0],arr[1])
	for i := 2;i < l;i++ {
		a := opt[i-2] + arr[i]
		b := opt[i-1]
		opt[i]=max(a,b)
	}
	return opt[l-1]
}


func max(a,b int)int{
	if a > b {
		return a
	}
	return b
}