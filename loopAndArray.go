package main

import (
	"fmt"
)

var numbers = []int{1, 2, 3, 4}

func main() {
	// numbers := []int{1, 2, 3, 5, 6}
	for i, num := range numbers {
		fmt.Printf("%d: %d\n", i, num)
	}
	fmt.Println("-=-=-=-=-=-=-=-=-=-=-=---=-=")
	valu()
}

func valu() {
	for i := 0; i < len(numbers); i++ {
		fmt.Println(numbers[i])
	}
	fmt.Println("-=-=-=-==-=-=-=-=-=-=-")
	inhance()
}
func inhance() {
	for _, num := range numbers {
		fmt.Println(num)
	}
}
