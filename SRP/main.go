package main

import (
	"fmt"
	"time"
)

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

	manager := NewLogManager("logs.txt")
	manager.StartWorkers(3)

	for i := 1; i <= 10; i++ {
		manager.Log(fmt.Sprintf("Event #%d happened", i))
		time.Sleep(30 * time.Millisecond)
	}

	manager.Stop()
}
