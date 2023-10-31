package main

import "fmt"

type Speaker interface {
	Speak()
	Laugh()
}

type employee struct {
	name string
	age  int
}

func (e employee) Speak() {
	fmt.Println("Hello I am Mohammad ", e.name)
}
func (e employee) Laugh() {
	fmt.Println("Laughing ")
}
func SaySomeThing(s Speaker) {
	s.Speak()
	s.Laugh()
}

func main() {
	emp := employee{name: "Rizwan", age: 26}
	SaySomeThing(emp)
}
