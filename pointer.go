package main

import "fmt"

type student struct {
	name string
	age  int
}

func (s *student) birthday() {
	s.age = s.age + 1
}

func main() {
	s1 := student{
		name: "Priyanshu",
		age:  21,
	}
	fmt.Println("Before:", s1.age)
	s1.birthday()
	fmt.Println("After:", s1.age)
}

// Value method → s = copy
// Pointer method → s = location of object
// Method gets memory address of object
// Not copy.