package sumEqual

import (
	"fmt"
	"testing"
)

var arr = []int{3,34,4,12,5,2}

func TestRec_sumEqual(t *testing.T) {
	fmt.Println("<-----------------  Rec_sumEqual ------------------->")
	fmt.Println(Rec_sumEqual(arr,len(arr)-1,9))
	fmt.Println(Rec_sumEqual(arr,len(arr)-1,10))
	fmt.Println(Rec_sumEqual(arr,len(arr)-1,11))
	fmt.Println(Rec_sumEqual(arr,len(arr)-1,12))
	fmt.Println(Rec_sumEqual(arr,len(arr)-1,13))
}

func TestDp_sumEqual(t *testing.T) {
	fmt.Println("<-----------------  Dp_sumEqual ------------------->")
	fmt.Println(Dp_sumEqual(arr,9))
	fmt.Println(Dp_sumEqual(arr,10))
	fmt.Println(Dp_sumEqual(arr,11))
	fmt.Println(Dp_sumEqual(arr,12))
	fmt.Println(Dp_sumEqual(arr,13))
}