package config

import (
	"log"
	"os"

	"github.com/spf13/viper"
)

func ProvideConfig(path string) func() (Configuration, error) {
	return func() (Configuration, error) {
		return LoadConfig(path)
	}
}

type Configuration struct {
	AppName string `mapstructure:"APP_NAME"`
	AppEnv  string `mapstructure:"APP_ENV"`
	AppPort string `mapstructure:"APP_PORT"`
	DBHost  string `mapstructure:"DB_HOST"`
	DBPort  string `mapstructure:"DB_PORT"`
	DBUser  string `mapstructure:"DB_USER"`
	DBPass  string `mapstructure:"DB_PASS"`
	DBName  string `mapstructure:"DB_NAME"`
}

func LoadConfig(path string) (config Configuration, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName(".env")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	if err = viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			log.Println("No .env file found, reading from OS environment variables instead")
		} else {
			log.Fatalf("Fatal error reading config: %s", err)
		}
	}

	err = viper.Unmarshal(&config)

	// Use OS environment variables if the .env is missing some values
	if config.AppName == "" {
		config.AppName = os.Getenv("APP_NAME")
	}
	if config.AppEnv == "" {
		config.AppEnv = os.Getenv("APP_ENV")
	}
	if config.AppPort == "" {
		config.AppPort = os.Getenv("APP_PORT")
	}
	if config.DBHost == "" {
		config.DBHost = os.Getenv("DB_HOST")
	}
	if config.DBPort == "" {
		config.DBPort = os.Getenv("DB_PORT")
	}
	if config.DBUser == "" {
		config.DBUser = os.Getenv("DB_USER")
	}
	if config.DBPass == "" {
		config.DBPass = os.Getenv("DB_PASS")
	}
	if config.DBName == "" {
		config.DBName = os.Getenv("DB_NAME")
	}

	return config, nil
}
