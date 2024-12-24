package db

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"fiberProject/Fiber3/models" // Update the import path based on your project structure
)

var dbClient *gorm.DB

// Connect initializes the database connection
func Connect() {
	var err error
	// Connection string for the database
	dsn := "username:Password@tcp(localhost:3306)/Database?charset=utf8mb4&parseTime=True&loc=Local"
	dbClient, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Error connecting to the database: ", err)
	}

	// Test the underlying SQL database connection
	sqlDB, err := dbClient.DB()
	if err != nil {
		log.Fatal("Error getting SQL DB from GORM: ", err)
	}
	err = sqlDB.Ping()
	if err != nil {
		CloseDB()
		log.Fatal("Error pinging the database: ", err)
	}

	log.Println("Database connected")
}

// AutoMigrate runs database migrations
func AutoMigrate() {
	err := dbClient.AutoMigrate(&models.User{}) // Struct name should be exported (User, not user)
	if err != nil {
		log.Fatal("Error during migration: ", err)
	}
	log.Println("Database migration completed")
}

// GetDB provides the database client instance
func GetDB() *gorm.DB {
	return dbClient
}

// CloseDB closes the database connection
func CloseDB() {
	sqlDB, err := dbClient.DB()
	if err != nil {
		log.Println("Error getting SQL DB for closure: ", err)
		return
	}
	err = sqlDB.Close()
	if err != nil {
		log.Println("Error closing the database: ", err)
	} else {
		log.Println("Database connection closed")
	}
}
