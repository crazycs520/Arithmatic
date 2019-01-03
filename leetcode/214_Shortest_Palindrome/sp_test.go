package sp

import (
	"testing"
)

var cases = []struct {
	in  string
	out string
}{
	{"", ""},
	{"a", "a"},
	{"aa", "aa"},
	{"ab", "bab"},
	{"aab", "baab"},
	{"aacecaaa", "aaacecaaa"},
	{"abcd", "dcbabcd"},
	{"abb", "bbabb"},
}

func TestShortestPalindrome(t *testing.T) {
	for _, c := range cases {
		re := shortestPalindrome(c.in)
		if re != c.out {
			t.Fatalf("in: %v, out: %v, expect: %v\n", c.in, re, c.out)
		}
	}
}

func TestShortestPalindrome2(t *testing.T) {
	for _, c := range cases {
		re := shortestPalindrome2(c.in)
		if re != c.out {
			t.Fatalf("in: %v, out: %v, expect: %v\n", c.in, re, c.out)
		}
	}
}
