package main

import "fmt"

func ciaAgent() {
	name := "Priyanshu"
	age := 22
	city := "seattle"
	isStudent := false
	fmt.Printf(
		"Name : %s\nAge : %d\nCity : %s\nisStudent : %t\n",
		name,
		age,
		city,
		isStudent,
	)
}

func main() {
	ciaAgent()
}
