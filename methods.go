package main

import "fmt"



type Person struct {
	name string
	age  int
}
func (p Person) speek() {
	fmt.Println("Hello ", p.name)

}
func (p *Person) speeks() {
	fmt.Println("Hello Mohammad ", p.name)
}


func main() {
	person := Person{name: "Rizwan", age: 26}
	person.speek()
	person.speeks()
}
