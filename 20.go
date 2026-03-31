package main

import "fmt"

func main() {
	var a [2]string
	a[0] = "CAPTAIN"
	a[1] = "AMERICA"
	fmt.Println(a)
	fmt.Println(a[0], a[1])
	fmt.Println(a[1])

	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)
}
