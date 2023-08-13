package database

import (
	"fmt"
	"strconv"

	"github.com/nunenuh/iquote-fiber/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDB() {
	var err error
	p := config.Config("DB_PORT")
	port, err := strconv.ParseUint(p, 10, 32)

	if err != nil {
		panic("Failed to parse database port")
	}

	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		config.Config("DB_HOST"),
		port, config.Config("DB_USER"),
		config.Config("DB_PASSWORD"),
		config.Config("DB_NAME"))
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to database")
	}

	// fmt.Println("Connection Opened to Database")
	// DB.AutoMigrate(&model.Product{}, &model.User{})
	// fmt.Println("Database migrated!")

}
