package db

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/Turut4/GradeFlow/internal/store"
	"gorm.io/gorm"
)

var usernames = []string{
	"Bob", "Carlos", "Pietra", "Julia", "Miguel", "Sofia", "Lucas", "Ana",
	"Maria", "Jhon", "Laura", "Gabriel", "Isabella", "Rafael", "Amanda",
	"Diego", "Felipe", "Helena", "Eduardo", "Mariana", "Bruno", "Camila",
	"Ricardo", "Larissa", "Thiago", "Alice", "Daniel", "Luiza", "Rodrigo",
	"Paul", "Clara", "Andre", "Caroline", "Gustavo", "Isabela", "Victor",
	"Alex", "Marcel", "Bianca", "Matheus", "Renata",
}

func Seed(store store.Storage, db *gorm.DB) {
	users := generateUsers(300)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	for _, user := range users {
		if err := store.Users.Create(ctx, user); err != nil {
			log.Println("Error creating user:", err)
			return
		}
	}
}

func generateUsers(num int) []*store.User {
	users := make([]*store.User, num)

	for i := 0; i < num; i++ {
		users[i] = &store.User{
			Username: usernames[i%len(usernames)] + fmt.Sprintf("%d", i),
			Email:    usernames[i%len(usernames)] + fmt.Sprintf("%d", i) + "@example.com",
			Role: store.Role{
				Name: "user",
			},
		}
	}

	return users
}
