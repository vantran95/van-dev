package main

import (
	"fmt"
	"log"

	"hexagonal/adapters/repositories"
	"hexagonal/adapters/services"
	"hexagonal/core/domain"
)

// Cấu trúc thư mục:
// hexagonal-user-management/
// ├── core/
// │   ├── domain/
// │   │   └── user.go
// │   └── ports/
// │       ├── user_repository.go
// │       └── user_service.go
// ├── adapters/
// │   ├── repositories/
// │   │   └── in_memory_user_repository.go
// │   └── services/
// │       └── user_service.go
// └── main.go

func main() {
	repo := repositories.NewInMemoryUserRepository()
	userService := services.NewUserService(repo)

	user := &domain.User{
		Name:  "John Doe",
		Email: "john@example.com",
	}

	if err := userService.CreateUser(user); err != nil {
		log.Fatal(err)
	}

	users, err := userService.ListUsers()
	if err != nil {
		log.Fatal(err)
	}

	for _, u := range users {
		fmt.Printf("User: %s, Email: %s\n", u.Name, u.Email)
	}
}
