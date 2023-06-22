package database

import (
	"fmt"

	"github.com/michaelcychan/art-notebook/backend-go/config"
	"github.com/michaelcychan/art-notebook/backend-go/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Database instance
type Dbinstance struct {
	DB *gorm.DB
}

var DB Dbinstance

// Connect function
func Connect() (*gorm.DB, error) {

	db, err := gorm.Open(postgres.Open(config.Config("DB_URL")), &gorm.Config{})

	if err != nil {
		fmt.Println("Database connection FAILED")
		return db, err
	}

	fmt.Println("Database connection success")

	db.AutoMigrate(&models.SavedData{})
	return db, nil

}
