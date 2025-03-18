package db

import (
	"errors"
	"fmt"
	"log"

	"github.com/Turut4/GradeFlow/internal/store"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func NewDB(addr string) (*gorm.DB, error) {
	if addr == "" {
		return nil, errors.New("DSN cannot be empty")
	}

	config := &gorm.Config{
		Logger:      logger.Default.LogMode(logger.Info),
		PrepareStmt: true,
	}

	db, err := gorm.Open(postgres.Open(addr), config)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, fmt.Errorf("failed to get database instance: %w", err)
	}
	if err := sqlDB.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	if err := migrate(db); err != nil {
		return nil, fmt.Errorf("failed to run migrations: %w", err)
	}

	return db, nil
}

func migrate(db *gorm.DB) error {
	models := []any{
		&store.User{},
		&store.Role{},
	}

	for _, model := range models {
		if err := db.AutoMigrate(model); err != nil {
			return fmt.Errorf("failed to migrate model %T: %w", model, err)
		}
	}

	SeedRoles(db)

	return nil
}

func SeedRoles(db *gorm.DB) {
	roles := []store.Role{
		{Name: "admin", Description: "Administrator", Level: 3},
		{Name: "user", Description: "Regular User", Level: 2},
		{Name: "guest", Description: "Guest User", Level: 1},
	}

	for _, role := range roles {
		var count int64
		if err := db.Model(&store.Role{}).Where("name = ?", role.Name).Count(&count).Error; err != nil {
			log.Println("Erro ao verificar role:", err)
			continue
		}

		if count == 0 {
			if err := db.Create(&role).Error; err != nil {
				log.Println("Erro ao criar role:", err)
			}
		}
	}
}
