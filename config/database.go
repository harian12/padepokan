package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"padepokan/helpers"
	"padepokan/logger"
)

func ConnectionDB() *gorm.DB {
	err := godotenv.Load()
	if err != nil {
		logger.ErrorFunc(err.Error(), err)
		log.Fatal("Error Load ENV File")
	}

	DB_HOST := helpers.GetEnv("DB_HOST", "localhost")
	DB_PORT := helpers.GetEnv("DB_PORT", "3306")
	DB_DATABASE := helpers.GetEnv("DB_DATABASE", "kiosana")
	DB_USERNAME := helpers.GetEnv("DB_USERNAME", "root")
	DB_PASSWORD := helpers.GetEnv("DB_PASSWORD", "")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", DB_USERNAME, DB_PASSWORD, DB_HOST, DB_PORT, DB_DATABASE)
	db, errDB := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if errDB != nil {
		logger.ErrorFunc(err.Error(), err)

		panic("gagal koneksi ke Database")
	}
	return db
}
