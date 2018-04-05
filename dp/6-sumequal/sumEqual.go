package sumEqual

func Rec_sumEqual(arr []int,i, s int) bool{
	if i == 0 {
		return arr[i] == s
	}
	if s == 0 {
		return true
	}
	if arr[i] > s{
		return Rec_sumEqual(arr,i-1,s)
	}
	a := Rec_sumEqual(arr,i-1,s-arr[i])
	b := Rec_sumEqual(arr,i-1,s)
	return a || b
}

func Dp_sumEqual(arr []int , s int) bool{
	l := len(arr)
	opt := make([][]bool,l)
	for i:=0;i<l;i++{
		opt[i] = make([]bool,s+1)
	}

	for i:=0;i<s+1;i++{
		opt[0][i] = false
	}
	if arr[0] <= s {
		opt[0][arr[0]] = true
	}

	for i:=0;i<l;i++{
		opt[i][0]=true
	}
	for i :=1;i<l;i++{
		for j:=1;j<s+1;j++{
			if arr[i] > j{
				opt[i][j]=opt[i-1][j]
			}else{
				opt[i][j] = opt[i-1][j-arr[i]] || opt[i-1][j]
			}

		}
	}
	return opt[l-1][s]
}