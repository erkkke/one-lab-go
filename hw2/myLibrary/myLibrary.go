package myLibrary

func LowerToUpper(s string) string {
	result := ""
	for i := 0; i < len(s); i++ {
		ch := s[i]
		if 'a' <= ch && ch <= 'z' {
			ch -= 'a' - 'A'
		}
		result += string(ch)
	}
	return result
}

func UpperToLower(s string) string {
	result := ""
	for i := 0; i < len(s); i++ {
		ch := s[i]
		if 'A' <= ch && ch <= 'Z' {
			ch += 'a' - 'A'
		}
		result += string(ch)
	}
	return result
}