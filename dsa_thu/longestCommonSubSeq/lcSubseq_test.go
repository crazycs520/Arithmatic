package longestCommonSubSeq

import (
	"fmt"
	"testing"
)

func TestLCSubseq(t *testing.T) {
	a := "acgbfhkhytfrtu"
	b := "cegefkhuhweqtuj"
	s := LCSubseq(a, b)
	fmt.Println(s)
}

func TestLCSubseq1(t *testing.T) {
	a := "acgbfhkhytfrtuwertyuo"
	b := "cegefkhuhweqtujwertuop"
	s := LCSubseq1(a, b)
	fmt.Println(s)
}
