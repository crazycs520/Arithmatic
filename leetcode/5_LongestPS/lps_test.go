package lps

import (
	"testing"
)

type cases struct {
	in, out string
}

var cases1 = []cases{
	{"babad", "aba"},
	{"asdfggfdsa", "asdfggfdsa"},
	{"qweasdfggfdsauioh", "asdfggfdsa"},
}

func TestLps(t *testing.T) {
	for _, c := range cases1 {
		got := Lps(c.in)
		if got != c.out {
			t.Errorf("Lps(%s)=%s,but want %s\n", c.in, got, c.out)
		} else {
			t.Logf("Lps(%s)=%s, want %s\n", c.in, got, c.out)
		}
	}
}

func TestReverse(t *testing.T) {
	str := "abcdefg"
	t.Logf("reverse(%s)=%s   str=%s", str, Reverse(str), str)
}
