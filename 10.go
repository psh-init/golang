package main

import "fmt"

func main() {
	names := []string{"Priyanshu", "Tony stark", "Bruce wayne", "Thor", "Loki "}
	fmt.Println(names)
	fmt.Println(names[2])
	fmt.Println(len(names))
	names = append(names, "Wanda")
	names = append(names, "Vision")
	names = append(names, "Captain America")
	fmt.Println(names)
	for _, names := range names {
		fmt.Println(names)
	}

}
