package replaceBlank

import (
	"fmt"
	"testing"
)

func TestReplaceBlank(t *testing.T) {
	s1 := "hello crazycs ."
	ss := ReplaceBlank(s1)
	fmt.Println(ss)
	s1 = " crazycs "
	ss = ReplaceBlank(s1)
	fmt.Println(ss)
	s1 = " "
	ss = ReplaceBlank(s1)
	fmt.Println(ss)
	s1 = ""
	ss = ReplaceBlank(s1)
	fmt.Println(ss)
}
