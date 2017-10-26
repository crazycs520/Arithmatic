package lps

func Lps(str string) string {
	str1 := Reverse(str)
	end := 0
	l, ll := 0, 0
	i := 0
	for i = 0; i < len(str); i++ {
		if str[i] == str1[i] {
			l++
			end = i
		} else {
			if l > ll {
				ll = l
				end = i
			}
			l = 0

		}
	}
	if l > ll {
		ll = l
		end = i
	}
	return str[end-ll : end]
}

func Reverse(str string) string {
	s := []byte{}
	for i := len(str) - 1; i >= 0; i-- {
		s = append(s, str[i])
	}
	return string(s)
}
