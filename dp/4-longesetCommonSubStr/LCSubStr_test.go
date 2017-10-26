package dp

import (
	"fmt"
	"testing"
)

func TestLCSubStr(t *testing.T) {
	s1 := "YXXXXXXYXU"
	s2 := "YXYXXYYYYXXYYYYXYYXXYYXXYXYYYYYYXYXYYXYXYYYXXXXXXYXU"

	s, l := LCSubStr(s1, s2)
	fmt.Println(s, len(s), l)

}
