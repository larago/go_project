package main

import "fmt"

func main() {
	var a int = 20
	var ip *int

	a = 30

	ip = &a

	fmt.Printf("a address is: %x\n", &a)

	fmt.Printf("ip address is: %x\n", ip)

	fmt.Printf("ip value is %d\n", *ip)
}
