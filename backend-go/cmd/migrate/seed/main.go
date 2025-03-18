package main

import (
	"log"

	"github.com/Turut4/GradeFlow/internal/db"
	"github.com/Turut4/GradeFlow/internal/env"
	"github.com/Turut4/GradeFlow/internal/store"
)

func main() {
	addr := env.GetString("DB_ADDR", "postgres://user:password@localhost:5432/gradeflow?sslmode=disable")
	conn, err := db.NewDB(addr)
	if err != nil {
		log.Fatal(err)
	}

	store := store.NewStorage(conn)
	db.Seed(store, conn)
}
