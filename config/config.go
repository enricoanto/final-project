package config

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/spf13/viper"

	_ "github.com/lib/pq"
)

var (
	DB *gorm.DB
)

type DBConfig struct {
	Host     string
	Port     string
	User     string
	DBName   string
	Password string
}

func GetDBConfig() *DBConfig {
	return &DBConfig{
		Host:     viper.GetString("DB_HOST"),
		Port:     viper.GetString("DB_PORT"),
		User:     viper.GetString("DB_USER"),
		Password: viper.GetString("DB_PASSWORD"),
		DBName:   viper.GetString("DB_NAME"),
	}
}

func init() {
	loadEnv()
	databaseConnect()
}

func (config *DBConfig) GetDBUrl() string {
	return fmt.Sprintf("user=%v password=%v host=%v sslmode=disable", config.User, config.Password, config.Host)
}

func databaseConnect() {
	dsn := GetDBConfig().GetDBUrl() + fmt.Sprintf(" dbname=%v", GetDBConfig().DBName)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("error connecting DB:", err)
	}

	DB = db
}

func loadEnv() {
	viper.SetConfigFile(".env")
	err := viper.ReadInConfig()

	if err != nil {
		log.Fatalf("Error while reading config file %s", err)
	}
}
