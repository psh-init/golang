package main

import "fmt"

func PrintAvengers(avengers []string) {
	for i, name := range avengers {
		fmt.Printf("%d %s\n", i+1, name)
	}
}

func main() {
	avengers := []string{"THOR", "HULK", "NATASHA"}
	PrintAvengers(avengers)
}
