package main

import "fmt"

func userInfo(name string, age int) string {
	return fmt.Sprintf("%s is %d years old.", name, age)
}

func main() {
	result := userInfo("Priyanshu", 22)
	fmt.Println(result)
}
