package main

import "fmt"

type student struct {
	name string
	age  int
}

func (s student) intro() {
	fmt.Println("hii My name is", s.name)
}

func main() {
	s1 := student{
		name: "Priyanshu",
		age:  21,
	}
	s1.intro()
}
