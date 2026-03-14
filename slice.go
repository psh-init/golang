package main

import "fmt"

func main() {
	tasks := []string{}
	tasks = append(tasks, "EAT")
	tasks = append(tasks, "CODE")
	tasks = append(tasks, "REPEAT")
	fmt.Println("YOUR TASKS:")
	for i := 0; i < len(tasks); i++ {
		fmt.Println(i+1, "-", tasks[i])
	}
}
