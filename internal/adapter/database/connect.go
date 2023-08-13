package database

import (
	"fmt"
	"time"

	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connection() (db *gorm.DB) {

	host := viper.Get("DB_HOST")
	user := viper.Get("DB_USER")
	pass := viper.Get("DB_PASS")
	dbname := viper.Get("DB_NAME")
	dsn := fmt.Sprintf("host=%v user=%v password=%v dbname=%v sslmode=disable",
		host, user, pass, dbname)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	sqlDB, err := db.DB()
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)
	return db
}
