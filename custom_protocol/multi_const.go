package main

import "fmt"

func main() {
	const Length int = 10
	const Width int = 5
	var area int
	const a, b, c = 1, false, "str"

	area = Length * Width
	fmt.Printf("area is %d", area)
	println()
	println(a, b, c)
}
