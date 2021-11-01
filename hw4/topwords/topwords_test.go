package topwords

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTopWords(t *testing.T) {
	str1 := "hello, my name is Jack and nice to meet u. hello Jack."
	n1 := 2
	expected1 := []string{"hello", "jack"}
	result1 := TopWords(str1, n1)

	str2 := "abcd bbhasd kmxvkmzc qwerty ajsndasdz asdn njsd jnsdbasd jasjdn abcd asd abcd nnbasd abcd asd qwerty jas qwerty."
	n2 := 3
	expected2 := []string{"abcd", "asd", "qwerty"}
	result2 := TopWords(str2, n2)

	assert.Equal(t, expected1, result1)
	assert.Equal(t, expected2, result2)
}
