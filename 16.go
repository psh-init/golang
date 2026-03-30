package main

import "fmt"

type User struct {
	Name   string
	Age    int
	Mobile string
}

func main() {
	u := User{
		Name:   "Tony Stark",
		Age:    48,
		Mobile: "6969696969",
	}
	fmt.Println(u)
	fmt.Println(u.Name)
}
