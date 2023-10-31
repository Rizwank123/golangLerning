package main

import (
	"fmt"
)

// define struct
type Ractangle struct {
	width  int
	height int
}

// method to calculate area
func (r Ractangle) Area() int {
	return r.width * r.height
}

// method to calculate Perimeter
func (r Ractangle) Perimeter() int {
	return 2 * (r.width + r.height)
}

func main() {
	//create object of Struct
	var width, height int

	//input:=bufio.Reader(os.Stdin);
	fmt.Println("Enter the width: ")
	fmt.Scanf("%d", &width)
	fmt.Println("Enter the height: ")
	fmt.Scanf("%d", &height)

	rect := Ractangle{width: width, height: height}
	fmt.Println("Area: ", rect.Area())
	fmt.Println("Perimeter:", rect.Perimeter())
}
