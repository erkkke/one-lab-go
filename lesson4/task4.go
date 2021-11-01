package main

import (
	"fmt"
	"reflect"
	"strings"
)

// countRunes - TODO: implement
func countRunes(s string) map[rune]int {
	res := make(map[rune]int)
	for _, v := range []rune(s) {
		res[v] += 1
	}

	return res
}

func main() {
	s := "каждый охотник желает знать, где сидит фазан"
	expected := map[rune]int{
		'г': 1,
		'с': 1,
		'о': 2,
		'л': 1,
		'ь': 1,
		',': 1,
		'д': 3,
		'т': 4,
		'и': 3,
		'ы': 1,
		'х': 1,
		'з': 2,
		' ': 6,
		'н': 3,
		'е': 3,
		'ф': 1,
		'к': 2,
		'а': 5,
		'ж': 2,
		'й': 1,
	}
	got := countRunes(s)
	if reflect.DeepEqual(got, expected) {
		fmt.Println("OK")
	} else {
		fmt.Printf("expected:\n%s\ngot:\n%s\n", prettyFormat(expected), prettyFormat(got))
	}
}

func prettyFormat(mp map[rune]int) string {
	var result strings.Builder
	result.WriteString("{\n")
	var lines []string
	for r, count := range mp {
		lines = append(lines, fmt.Sprintf("\t'%c': %d", r, count))
	}
	result.WriteString(strings.Join(lines, ",\n"))
	result.WriteString("\n}")
	return result.String()
}
