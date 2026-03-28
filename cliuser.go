package main

import (
	"log"
	"os"
)

// Model
type User struct {
	Name string
}

// In-memory storage
var users []User

// Add user
func addUser(name string) {
	user := User{Name: name}
	users = append(users, user)
	log.Println("User added:", name)
}

// List users
func listUsers() {
	if len(users) == 0 {
		log.Println("No users found")
		return
	}

	log.Println("Users:")
	for _, user := range users {
		log.Println("-", user.Name)
	}
}

func main() {
	args := os.Args

	if len(args) < 2 {
		log.Println("No command provided")
		return
	}

	command := args[1]

	switch command {

	case "add":
		if len(args) < 3 {
			log.Println("No name provided")
			return
		}
		addUser(args[2])

	case "list":
		listUsers()

	default:
		log.Println("Invalid command")
	}
}
