//最长不重复字符子串
//Longest Substring Without Repeating Characters
package main

import (
	"fmt"
)

//1, Brute Force
func lengthOfLongestSubstring(s string) int {
	l, ll := 0, 0
	strlen := len(s)
	for i, j := 0, 0; i < strlen; i++ {
		chars := make(map[byte]bool)
		for j = i; j < strlen; j++ {
			if chars[s[j]] == false {
				l++
				chars[s[j]] = true
			} else {
				break
			}
		}
		if l > ll {
			ll = l
		}
		l = 0
	}
	return ll
}

//2, Sliding Window
func lengthOfLongestSubstring2(s string) int {
	l := 0
	strlen := len(s)
	chars := make(map[byte]bool)
	i, j := 0, 0
	for i < strlen && j < strlen {
		if chars[s[j]] == false {
			chars[s[j]] = true
			j++
			if (j - i) > l {
				l = j - i
			}
		} else {
			delete(chars, s[i])
			i++
		}
	}
	return l
}

//#3 Sliding Window Optimized [Accepted]
func lengthOfLongestSubstring3(s string) int {
	l := 0
	strlen := len(s)
	chars := make(map[byte]int)
	i, j := 0, 0
	for ; j < strlen; j++ {
		if _, ok := chars[s[j]]; ok {
			i = chars[s[j]]
		}
		if (j - i + 1) > l {
			l = j - i + 1
		}
		chars[s[j]] = j + 1

	}
	return l
}

//#4 (Assuming ASCII 128)
/*
Commonly used tables are:

int[26] for Letters 'a' - 'z' or 'A' - 'Z'
int[128] for ASCII
int[256] for Extended ASCII
*/

func lengthOfLongestSubstring4(s string) int {
	l := 0
	strlen := len(s)
	var chars [128]int
	i, j := 0, 0
	for ; j < strlen; j++ {
		if chars[s[j]] > i {
			i = chars[s[j]]
		}
		if (j - i + 1) > l {
			l = j - i + 1
		}
		chars[s[j]] = j + 1

	}
	return l
}

func main() {
	str := "abdfa"
	fmt.Println(lengthOfLongestSubstring(str))
	fmt.Println(lengthOfLongestSubstring2(str))
	fmt.Println(lengthOfLongestSubstring3(str))
	fmt.Println(lengthOfLongestSubstring4(str))
}
