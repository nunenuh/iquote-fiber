package database

import (
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/nunenuh/iquote-fiber/internal/database/model"
	"github.com/nunenuh/iquote-fiber/pkg/config"
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

	// db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
	// 	Logger: logger.New(
	// 		log.New(os.Stdout, "\r\n", log.LstdFlags), // You can replace os.Stdout with a log file if desired
	// 		logger.Config{
	// 			SlowThreshold: time.Second, // Set the threshold for slow query logging
	// 			LogLevel:      logger.Info, // Log level for the logger
	// 			Colorful:      true,        // Set to true to enable colorful output
	// 		},
	// 	),
	// })
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	maxOpenConn, err := strconv.Atoi(config.DBMaxIdleConns)
	maxIdleConn, err := strconv.Atoi(config.DBMaxOpenConns)

	sqlDB, err := db.DB()
	sqlDB.SetMaxIdleConns(maxOpenConn)
	sqlDB.SetMaxOpenConns(maxIdleConn)
	sqlDB.SetConnMaxLifetime(time.Hour)
	db.AutoMigrate(&model.User{}, &model.Author{}, &model.Category{}, &model.Quote{})
	return db, err
}
