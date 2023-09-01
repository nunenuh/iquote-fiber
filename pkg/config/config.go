package config

import (
	"log"
	"os"
	"regexp"
	"strconv"
	"time"

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
	AppHost string `mapstructure:"APP_HOST"`
	AppPort string `mapstructure:"APP_PORT"`

	DBHost         string `mapstructure:"DB_HOST"`
	DBPort         string `mapstructure:"DB_PORT"`
	DBUser         string `mapstructure:"DB_USER"`
	DBPass         string `mapstructure:"DB_PASS"`
	DBName         string `mapstructure:"DB_NAME"`
	DBMaxOpenConns string `mapstructure:"DB_MAX_OPEN_CONNS"`
	DBMaxIdleConns string `mapstructure:"DB_MAX_IDLE_CONNS"`
	// DBConnMaxLifetime time.Duration `mapstructure:"DB_CONN_MAX_LIFETIME"`

	JWTSecret    string `mapstructure:"JWT_SECRET"`
	JWTExpire    string `mapstructure:"JWT_EXPIRE"`
	JWTExpireInt int64
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

	setupConfig(&config)

	return config, nil
}

func setupConfig(config *Configuration) {
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

	if config.DBMaxOpenConns == "" {
		config.DBMaxOpenConns = os.Getenv("DB_MAX_OPEN_CONNS")
	}

	if config.DBMaxIdleConns == "" {
		config.DBMaxIdleConns = os.Getenv("DB_MAX_IDLE_CONNS")
	}

	if config.JWTSecret == "" {
		config.JWTSecret = os.Getenv("JWT_SECRET")
	}
	if config.JWTExpire == "" {
		config.JWTExpire = os.Getenv("JWT_EXPIRE")

		regex := regexp.MustCompile(`(\d+)([dhm])`)
		match := regex.FindStringSubmatch(config.JWTExpire)
		if len(match) == 3 {
			duration, err := strconv.Atoi(match[1])
			if err != nil {
				panic(err)
			}
			unit := match[2]

			switch unit {
			case "d":
				config.JWTExpireInt = time.Now().AddDate(0, 0, duration).Unix()
			case "h":
				config.JWTExpireInt = time.Now().Add(time.Hour * time.Duration(duration)).Unix()
			case "m":
				config.JWTExpireInt = time.Now().Add(time.Minute * time.Duration(duration)).Unix()
			default:
				// handle invalid unit
				panic("Invalid unit in JWTExpire value")
			}
		} else {
			// handle invalid JWTExpire value
			panic("Invalid JWTExpire value")
		}

		log.Printf("JWT_EXPIRE: %v", config.JWTExpire)
	}

}
