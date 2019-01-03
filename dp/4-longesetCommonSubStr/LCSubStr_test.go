package dp

import (
	"testing"
)

var testCase = []struct {
	s1 string
	s2 string
	r  string
}{
	{"", "", ""},
	{"a", "a", "a"},
	{"a", "b", ""},
	{"aa", "ba", "a"},
	{"abcde", "abcdg", "abcd"},
	{"abcde", "cdghabc", "abc"},
	{"abcdefg1234567890", "hijklmnopqrstuvwxyz0", "0"},
	{"abcdefghijklmnopqrstuvwxyz","abc0abcdef1cdefghijklmn","cdefghijklmn"},
}

func TestLCSubStr(t *testing.T) {
	for _, c := range testCase {
		re := LCSubStr(c.s1, c.s2)
		if re != c.r {
			t.Fatalf("%v, %v, expect %v, get: %v\n", c.s1, c.s2, c.r, re)
		}
	}
}

func TestLCSubStr2(t *testing.T) {
	for _, c := range testCase {
		re := LCSubStr2(c.s1, c.s2)
		if re != c.r {
			t.Fatalf("%v, %v, expect %v, get: %v\n", c.s1, c.s2, c.r, re)
		}
	}
}

func BenchmarkLCSubStr(b *testing.B) {
	for i:=0;i<b.N;i++ {
		for _, c := range testCase {
			LCSubStr(c.s1, c.s2)
		}
	}
}

func BenchmarkLCSubStr2(b *testing.B) {
	for i:=0;i<b.N;i++ {
		for _, c := range testCase {
			LCSubStr2(c.s1, c.s2)
		}
	}
}
