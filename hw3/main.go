package main

import (
	"fmt"
	"github.com/erkkke/one-lab-go/hw3/atoi"
	"github.com/erkkke/one-lab-go/hw3/fibonacci"
	"github.com/erkkke/one-lab-go/hw3/itoa"
	"github.com/erkkke/one-lab-go/hw3/reverse"
	"github.com/erkkke/one-lab-go/hw3/runeByIndex"
	"github.com/erkkke/one-lab-go/hw3/sort_imports"
	"log"
)

func main() {
	fmt.Println(itoa.Itoa(-22))
	fmt.Println(atoi.Atoi("144"))
	fmt.Println(reverse.Reverse("Еркебулан"))

	f := fibonacci.Fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Printf("%v ", f())
	}
	fmt.Println()

	s := "Еркебулан"
	ind := 6
	fmt.Println(runeByIndex.RuneByIndex(&s, &ind))

	if err := sort_imports.SortAndRewriteImports("sort_imports/test.go"); err != nil {
		log.Fatal(err)
	}

}


