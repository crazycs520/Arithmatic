package lisubseq

import (
	"fmt"
	"testing"
)

func TestLISubseq(t *testing.T) {
	a := []int{3, 5, 7, 1, 2, 8, 3, 4, 6}
	fmt.Println(LISubseq(a))
}
