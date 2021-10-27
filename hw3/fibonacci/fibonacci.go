package fibonacci

func Fibonacci() func() int {
	a, b := 0, 1
	return func() int {
		b, a = a, a + b
		return b
	}
}
