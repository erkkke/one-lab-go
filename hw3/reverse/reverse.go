package reverse

import "strings"

func Reverse(s string) string {
	res := new(strings.Builder)
	runeStr := []rune(s)
	for i := len(runeStr) - 1; i >= 0; i-- {
		res.WriteString(string(runeStr[i]))
	}
	return res.String()
}
