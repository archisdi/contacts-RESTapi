package models

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres" // what
	"github.com/joho/godotenv"
)

// database instance
var db *gorm.DB

func init() {
	e := godotenv.Load() // Load .env file
	if e != nil {
		fmt.Print(e)
	}

	username := os.Getenv("db_user")
	password := os.Getenv("db_pass")
	dbName := os.Getenv("db_name")
	dbHost := os.Getenv("db_host")
	dbPort := os.Getenv("db_port")

	// Build connection string
	dbURI := fmt.Sprintf("host=%s user=%s port=%s dbname=%s sslmode=disable password=%s", dbHost, username, dbPort, dbName, password)
	fmt.Println(dbURI)

	conn, err := gorm.Open("postgres", dbURI)
	if err != nil {
		fmt.Print(err)
	}

	db = conn
	db.Debug().AutoMigrate(&Account{}, &Contact{}) // Database migration
}

// GetDB returns a handle to the DB object (database instance)
func GetDB() *gorm.DB {
	return db
}
