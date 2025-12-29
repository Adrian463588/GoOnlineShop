package database

import (
	"ecommerce-backend/models"
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	// Configuration (In production, load from ENV)
	dsn := "root:@tcp(127.0.0.1:3306)/evermos_db?charset=utf8mb4&parseTime=True&loc=Local"

	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	fmt.Println("Database connected successfully")

	// Auto Migrate
	err = DB.AutoMigrate(
		&models.User{},
		&models.Address{},
		&models.Store{},
		&models.Category{},
		&models.Product{},
		&models.ProductPhoto{},
		&models.Transaction{},
		&models.TransactionDetail{},
		&models.ProductLog{},
	)
	if err != nil {
		log.Fatal("Failed to migrate database:", err)
	}
}