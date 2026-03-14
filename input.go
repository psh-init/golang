package main

import "fmt"

func main() {
	var name string
	var age int
	var city string
	fmt.Println("ENTER YOUR NAME:")
	fmt.Scanln(&name)
	fmt.Println("ENTER YOUR AGE:")
	fmt.Scanln(&age)
	fmt.Println("ENTER YOUR CITY")
	fmt.Scanln(&city)
	fmt.Println("\n----USER INFO----")
	fmt.Println("Name:", name)
	fmt.Println("Age:", age)
	fmt.Println("City:", city)
}
