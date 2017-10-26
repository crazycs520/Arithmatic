package fibonacci

func Fib1(n int) int {
	if n < 2 {
		return n
	}
	return Fib1(n-1) + Fib1(n-2)
}

func Fib2(n int) int {
	f, g := 0, 1
	for ; n > 0; n-- {
		g = g + f
		f = g - f
	}
	return f
}
