package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())
	n := 20
	a := make([]int, 0, n)
	for i := 0; i < n; i++ {
		a = append(a, rand.Intn(100))
	}
	fmt.Println(a)
	minInd := 0
	for i := range a {
		if a[minInd] > a[i] {
			minInd = i
		}
	}
	fmt.Println(a[minInd+1:])
}

