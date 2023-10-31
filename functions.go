package main

import (
	"fmt"
)

func main() {
	// var a, b int
	// fmt.Println("Enter value a and b using space seprated : ")
	// fmt.Scanf("%d %d", a, b)
	// fmt.Println("Enter value in b: ")
	// fmt.Scanf("%d", &b);
	sum, diff := calculateSum_Diff(14, 12)
	fmt.Printf("Sum is %d\n: ", sum)
	fmt.Printf("difference is : %d\n", diff)
	fmt.Println("variable argument function:\n", varadicFunction(1, 2, 3, 4, 5, 7, 8))
}
func calculateSum_Diff(a int, b int) (int, int) {
	return a + b, a - b
}

func varadicFunction(numbers ...int) int {
	sum := 0
	for number := range numbers {
		sum += number
	}
	return sum

}
