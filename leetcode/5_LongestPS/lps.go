package lps

func longestPalindrome(s string) string {
	if len(s) == 0 {
		return ""
	}
	i1, i2 := 0, len(s)
	start, end := 0, 1
	for i1 < i2 {
		if isPalindrome(s[i1:i2]) {
			if (i2 - i1) > (end - start) {
				start, end = i1, i2
			}
		}
		i2--
		if i2 == i1 {
			i1++
			i2 = len(s)
			if (i2 - i1) < (end - start) {
				break
			}
		}
	}
	return s[start:end]
}

func isPalindrome(s string) bool {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		if s[i] != s[j] {
			return false
		}
	}
	return true
}

func longestPalindrome2(s string) string {
	if len(s) == 0 {
		return ""
	}
	start,end := 0,0
	for i:=0;i<len(s);i++ {
		l1 :=  expandAroundCenter(s,i,i)
		l2 := expandAroundCenter(s,i,i+1)
		ll := max(l1,l2)
		if ll > (end - start){
			start = i - (ll -1)/2
			end  = i + ll /2
		}
	}
	return s[start:end+1]
}

func expandAroundCenter(s string, l,r int) int {
	for  ;l >= 0 && r < len(s) && s[l] == s[r]; {
		l--
		r++
	}
	return r - l - 1;
}

func max(a,b int) int {
	if a > b {
		return a
	}
	return b
}
