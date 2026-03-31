package main

import "fmt"

func main() {
	i, j := 42, 75
	p := &i         // points to i
	fmt.Println(*p) // read i through the pointer
	*p = 21         // set i through the pointer
	fmt.Println(i)  // see the new value of i

	p = &j // points to j
	fmt.Println(*p)
	*p = *p / 5    // divide j through the pointer
	fmt.Println(j) // see the new value of j
}
