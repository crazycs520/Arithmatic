package main

import (
	"fmt"
)

func main() {
	s := "ababg"
	fmt.Println(reverse(s))
	fmt.Println(longPa(s))
}

func longPa(s string) string {
	s1 := reverse(s)
	return long(s, s1)
}

func reverse(s string) string {
	b := []byte(s)
	for i, j := 0, len(s)-1; i < j; {
		b[i], b[j] = b[j], b[i]
		i++
		j--
	}
	return string(b)
}

func long(s1, s2 string) string {
	start, end := 0, 1
	for i := 0; i < len(s1); i++ {
		if (len(s1) - i) < (start - end) {
			break
		}
		for i1 := 0; i1 < len(s2); i1++ {
			if s1[i] == s2[i1] && (len(s2)-i1) > (end-start) {
				k, j := i, i1
				for k < len(s1) && j < len(s2) && s1[k] == s2[j] {
					k++
					j++
				}
				if (k - i) > (end - start) {
					start = i
					end = k
				}
			}
		}
	}
	return s1[start:end]
}
