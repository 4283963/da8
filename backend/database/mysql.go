package database

import (
	"fmt"
	"jade-grading/config"
	"jade-grading/models"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Init(cfg *config.DatabaseConfig) error {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.DBName)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return fmt.Errorf("failed to connect database: %w", err)
	}

	err = db.AutoMigrate(&models.JadeBracelet{})
	if err != nil {
		return fmt.Errorf("failed to auto migrate: %w", err)
	}

	DB = db
	log.Println("Database connection established successfully")
	return nil
}

func GetDB() *gorm.DB {
	return DB
}
