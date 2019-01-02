package lps

import (
	"fmt"
	"os"
	"testing"
)

var testCase = []struct {
	s string
	r string
}{
	{"babac", "bab"},
	{"cbbd", "bb"},
	{"a", "a"},
	{"aa", "aa"},
	{"", ""},
}

func TestLongestPalindrome(t *testing.T) {
	for _, c := range testCase {
		re := longestPalindrome(c.s)
		if re != c.r {
			fmt.Printf("%v, expect %v, get: %v\n", c.s, c.r, re)
			os.Exit(1)
		}
	}
}

func TestLongestPalindrome2(t *testing.T) {
	for _, c := range testCase {
		re := longestPalindrome2(c.s)
		if re != c.r {
			fmt.Printf("%v, expect %v, get: %v\n", c.s, c.r, re)
			os.Exit(1)
		}
	}
}
