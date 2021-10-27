package itoa

func Itoa(a int) string {
	res := ""
	isNegative := false

	if a < 0 {
		isNegative = true
		a *= -1
	}

	for a > 0 {
		tmp := a % 10
		a /= 10
		res = string(rune(tmp) + '0') + res
	}

	if isNegative {
		return "-" + res
	}
	return res
}
