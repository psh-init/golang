package main

import "fmt"

func printAvengers() {
	avengers := []string{"Priyanshu", "Tony stark", "Bruce wayne"}
	for i, avengers := range avengers {
		fmt.Printf("%d. %s\n", i+1, avengers)
	}
}

func main() {
	printAvengers()
}
