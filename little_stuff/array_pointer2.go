package main

import (
	"fmt"
)

const (
	Max int = 3
)

func main() {
	a := []int{10, 100, 200}
	var i int
	var ptr [Max]*int

	for i = 0; i < Max; i++ {
		ptr[i] = &a[i]
	}

	for i = 0; i < Max; i++ {
		fmt.Printf("a[%d] = %d\n", i, *ptr[i])
	}
}
