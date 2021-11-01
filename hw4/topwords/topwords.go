package topwords

import (
	"sort"
	"strings"
)

func TopWords(s string, n int) []string {
	s = strings.Replace(s, ",", "", -1)
	s = strings.Replace(s, ".", "", -1)
	words := strings.Split(strings.ToLower(s), " ")
	wordsCount := make(map[string]int)
	result := make([]string, n)

	for _, word := range words {
		wordsCount[word]++
	}

	for i := 0; i < n; i++ {
		maxi := 0
		for key, value := range wordsCount {
			if value > maxi {
				maxi = value
				result[i] = key
			}
		}
		delete(wordsCount, result[i])
	}
	sort.Strings(result)
	return result
}


