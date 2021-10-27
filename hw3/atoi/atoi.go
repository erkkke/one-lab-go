package atoi

import "errors"

var (
	errEmptyString = errors.New("error atoi, empty string")
	errInvalidSyntax = errors.New("error atoi, invalid syntax")
)

func Atoi(s string) (int, error) {
	if len(s) < 1 {
		return 0, errEmptyString
	}

	isNegative := false
	if s[0] == '-' || s[0] == '+' {
		if s[0] == '-' {
			isNegative = true
		}
		s = s[1:]
	}

	res := 0
	for _, ch := range []byte(s) {
		ch -= '0'
		if ch > 9 {
			return 0, errInvalidSyntax
		}
		res = res * 10 + int(ch)
	}
	if isNegative {
		res *= -1
	}
	return res, nil
}
