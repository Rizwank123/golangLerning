package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the text: ")
	text, _ := reader.ReadString('\n')
	fmt.Println("You enterd: ", text)
	var input string
	fmt.Printf("Enter the text: ")
	fmt.Scanln(&input)
	fmt.Println("You enterd: ", input)
	var name string
	var age int
	fmt.Println("enter your name ang age seprated by space: ")
	fmt.Scanf("%s %d", &name, &age)
	fmt.Printf("Name: %s")
}
