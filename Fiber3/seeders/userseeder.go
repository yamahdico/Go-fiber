package seeders

import (
	"log"

	db "fiberProject/Fiber3/database"
	"fiberProject/Fiber3/models"
)

func SeedUsers() {
	dbClient := db.GetDB()

	// Define a user
	user := models.User{
		Username: "UserTest1",
		Password: "1234567891",
		Email:    "user@domain.com1",
	}

	// Add the user to the database
	err := dbClient.Create(&user).Error
	if err != nil {
		log.Fatalf("User cannot be added to the database: %v", err)
	}

	log.Println("User seeded successfully")
}
