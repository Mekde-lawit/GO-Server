package utils

import (
	"fmt"

	"auth/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	host     = "localhost"
	port     = "5432"
	user     = "postgres"
	password = "pass1234"
	dbname   = "postgres"
)
// const (
// 	host     = "localhost"
// 	port     = "5433"
// 	user     = "postgres"
// 	password = "123456"
// 	dbname   = "crud"
// )

var db *gorm.DB

func InitDB() error {
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname,
	)
	
	var err error
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}

	err = db.AutoMigrate(&models.User{})
	if err != nil {
		return err
	}

	return nil
}

// GetDB simply returns the already initialized connection pool
func GetDB() *gorm.DB {
	if db == nil {
		panic("Database not initialized! Call InitDB first.")
	}
	return db
}