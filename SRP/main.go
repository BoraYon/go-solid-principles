package main

import "fmt"

type User struct {
	Name  string
	Email string
}

func printUser(user User) {
	fmt.Printf("Name: %s\nEmail: %s\n", user.Name, user.Email)
}

func main() {
	user := User{Name: "Bora", Email: "bora@example.com"}
	printUser(user)
}
