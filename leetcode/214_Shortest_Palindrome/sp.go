package sp

// 76ms
// 从中间开始，找能与第 0 个 char 构成最长的回文字符串
func shortestPalindrome(s string) string {
	if len(s) < 2 {
		return s
	}

	for mid := (len(s) - 1) / 2; mid >= 0; mid-- {
		c1 := can_palindrome(s, mid, mid+1)
		c2 := can_palindrome(s, mid-1, mid+1)
		if c1 {
			return finish_palindrom(s, mid, true)
		}
		if c2 {
			return finish_palindrom(s, mid, false)
		}
	}
	return finish_palindrom(s, 0, false)
}

func can_palindrome(s string, l, r int) bool {
	for l >= 0 && r < len(s) {
		if s[l] != s[r] {
			return false
		}
		l--
		r++
	}
	// think about "abb".
	if l >= 0 {
		return false
	}
	return true
}

func finish_palindrom(s string, mid int, isEven bool) string {
	ll := ((len(s) - 1) - mid) * 2
	if !isEven {
		ll++
	}
	bs := make([]byte, ll-len(s), ll)
	bs = append(bs, []byte(s)...)
	remain := ll - len(s)
	lastIdx := ll - 1
	for i := 0; i < remain; i++ {
		bs[i] = bs[lastIdx-i]
	}
	return string(bs)
}

//
func shortestPalindrome2(s string) string {
	i := 0
	for j := len(s) - 1; j >= 0; j-- {
		if s[i] == s[j] {
			i++
		}
	}
	if i == len(s) {
		return s
	}
	reverse_str := reverseString(s, i, len(s))
	return reverse_str + shortestPalindrome2(s[0:i]) + s[i:]
}

func reverseString(s string, l, r int) string {
	bs := make([]byte, 0, r-l)
	for i := r - 1; i >= l; i-- {
		bs = append(bs, s[i])
	}
	return string(bs)
}
