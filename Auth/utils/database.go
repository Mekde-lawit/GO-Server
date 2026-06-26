package utils 

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"auth/models"
)

const (
	host     = "localhost"
	port     = "5432"
	user     = "postgres"
	password = "pass1234"
	dbname   = "postgres"
)


func GetConnection() *gorm.DB{
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname,
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&models.User{})

	return db
}