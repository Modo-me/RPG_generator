package database

import (
	"fmt"
	"os"
	"rpg_generator/internal/module/generation/repository"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func DB_INIT() *gorm.DB {
	host := os.Getenv("DB_HOST")
	dsn := "user=usrnanme password=yourpasswd  host=%s port=5432 dbname=commentsDB sslmode=disable"
	dsn = fmt.Sprintf(dsn, host)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database")
	}

	err = DbSync(db)
	if err != nil {
		panic(err)
	}
	return db
}

func DbSync(db *gorm.DB) error {
	err := db.AutoMigrate(&repository.Story{})
	if err != nil {
		return err
	}
	return nil
}
