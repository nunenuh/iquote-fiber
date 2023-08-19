package database

import (
	"fmt"
	"time"

	"github.com/nunenuh/iquote-fiber/internal/adapter/config"
	"github.com/nunenuh/iquote-fiber/internal/adapter/database/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// ProvideDatabaseConnection now returns a function, not fx.Option
func ProvideDatabaseConnection() func(config config.Configuration) (*gorm.DB, error) {
	return Connection
}

func Connection(config config.Configuration) (db *gorm.DB, err error) {
	dsn := fmt.Sprintf("host=%v user=%v password=%v dbname=%v sslmode=disable",
		config.DBHost, config.DBUser, config.DBPass, config.DBName)
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	sqlDB, err := db.DB()
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)
	db.AutoMigrate(&model.User{}, &model.Author{}, &model.Category{}, &model.Quote{})
	return db, err
}
