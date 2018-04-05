package maxNJAS

import (
	"fmt"
	"testing"
)

var arr = []int{1,2,4,1,7,8,3}

func TestRec_maxSum(t *testing.T) {
	fmt.Println(Rec_maxSum(arr,len(arr)-1))
}

func TestDp_maxSum(t *testing.T) {
	fmt.Println(Dp_maxSum(arr))
}
