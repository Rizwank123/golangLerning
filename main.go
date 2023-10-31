package main

import "fmt"

func main() {
	for i := 0; i <= 10; i++ {
		a := 6
		fmt.Println(a * i)
		a = a * i
	}
}
func Hello(name string) string {
	return "Hello World"

}
