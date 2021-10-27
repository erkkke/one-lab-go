package runeByIndex

import "errors"

var (
	errNilIndex = errors.New("error RuneByIndex, nil index argument")
	errNilString = errors.New("error RuneByIndex, nil string argument")
	errInvalidIndex = errors.New("error RuneByIndex, invalid index value")
	errIndexOutOfRange = errors.New("error RuneByIndex, index out of range")
)

func RuneByIndex(s *string, i *int) (rune, error) {
	runes := []rune(*s)
	if s == nil {
		return 0, errNilString
	}
	if i == nil {
		return 0, errNilIndex
	}
	if *i < 0 {
		return 0, errInvalidIndex
	}
	if *i >= len(runes) {
		return 0, errIndexOutOfRange
	}
	return runes[*i], nil
}
