package main

import "fmt"

func loop() {
	for i := 1; i < 10; i++ {
		fmt.Println(i)
	}
}

func evenNumbers() {
	for i := 1; i <= 20; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}
}

func main() {
	loop()
	evenNumbers()
}
