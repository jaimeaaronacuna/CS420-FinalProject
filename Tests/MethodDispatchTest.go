package main

import "fmt"

type Speaker interface {
	Speak()
}

type Parent struct{}

func (Parent) Speak() {
	fmt.Println("Parent")
}

type Child struct{}

func (Child) Speak() {
	fmt.Println("Child")
}

func main() {
	var s Speaker

	s = Child{}

	s.Speak()
}

