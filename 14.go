package main

import "fmt"

func main() {
	user := map[string]string{
		"name":   "Priyanshu",
		"city":   "seattle",
		"age":    "22",
		"email":  "priyanshu@gmail.com",
		"number": "4437829",
	}
	fmt.Println(user)
	fmt.Println(user["name"])
	fmt.Println(user["email"])

}
