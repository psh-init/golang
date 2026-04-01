package main

import "fmt"

func main() {
	names := [4]string{
		"john",
		"harold",
		"ringo",
		"sher",
	}
	fmt.Println(names)

	a := names[0:2]
	b := names[1:3]
	fmt.Println(a, b)
	b[0] = "xxxx"
	fmt.Println(a, b)
	fmt.Println(names)

}
