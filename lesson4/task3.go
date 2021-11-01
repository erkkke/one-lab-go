package main

import (
	"fmt"
	"reflect"
)

func Flatten(slices [][]int) []int {
	res := make([]int, 0)
	for _, v := range slices {
		res = append(res, v...)
	}
	return res
}

func main() {
	type pair struct {
		s [][]int
		r []int
	}
	test := []pair{
		{[][]int{{1, 2}, {3, 4}}, []int{1, 2, 3, 4}},
		{[][]int{{1, 2}, {3, 4}, {6, 5}}, []int{1, 2, 3, 4, 6, 5}},
		{[][]int{{1, 2}, {}, {6, 5}}, []int{1, 2, 6, 5}},
	}
	for _, t := range test {
		s := t.s
		r := t.r
		r2 := Flatten(s)
		fmt.Printf("Test: %v\n", s)
		fmt.Printf("Expected: %v\n", r)
		fmt.Printf("Result: %v\n", r2)
		if reflect.DeepEqual(r, r2) {
			fmt.Printf("OK\n\n")
		} else {
			fmt.Printf("FAIL\n\n")
		}
	}
}

