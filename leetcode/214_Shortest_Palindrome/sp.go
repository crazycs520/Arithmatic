package sp

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
