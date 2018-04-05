package replaceBlank

func ReplaceBlank(s string) string {
	if len(s) < 1 {
		return ""
	}
	l := 0
	for _, v := range s {
		if v == ' ' {
			l++
		}
	}
	l = len(s) + l*2
	bs := make([]byte, l)
	i := 0
	for index := 0; index < len(s); index++ {
		if s[index] == ' ' {
			bs[i] = '%'
			bs[i+1] = '2'
			bs[i+2] = '0'
			i += 3
		} else {
			bs[i] = s[index]
			i++
		}
	}
	return string(bs)
}
