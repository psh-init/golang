package main

import "fmt"

func userProfile() {
	user := map[string]interface{}{
		"name":  "Priyanshu",
		"age":   22,
		"city":  "seattle",
		"email": "priyanshu@gmail.com",
	}
	fmt.Println(user)
	user["country"] = "USA"
	fmt.Println(user)
	delete(user, "email")
	fmt.Println(user)
}

func main() {
	userProfile()
}
