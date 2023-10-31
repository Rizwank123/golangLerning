package main

import "fmt"

func main() {
	a := 30
	b := &a
	fmt.Printf("%d\n", a)
	fmt.Printf("%d\n", b)
	fmt.Printf("%d\n", *b)
	*b = 45
	fmt.Printf("%d\n", a);
}
