package main

import "fmt"

func main() {
	var a int = 100
	var b int = 200

	fmt.Printf("Before exchange, a value is %d\n", a)
	fmt.Printf("Before exchange, b value is %d\n", b)

	swap(&a, &b)

	fmt.Printf("After exchanged, a value is %d\n", a)
	fmt.Printf("After exchanged, b value is %d\n", b)
}

func swap(x *int, y *int) {
	var tmp int
	tmp = *x
	*x = *y
	*y = tmp
}
