package initializers

import (
	"fmt"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error

func ConnectToDB() {
	// Get the database connection string from the environment variable
	dsn := os.Getenv("DB")
	if dsn == "" {
		panic("DB environment variable is not set")
	}

	// Open a connection to the MySQL database
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("Failed to connect to the database:", err)
		panic("Cannot connect to database")
	}

	// Print a message indicating successful connection
	fmt.Println("Connected to the database")
}
