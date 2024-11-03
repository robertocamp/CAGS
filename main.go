package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"cags/api/routes"

	_ "github.com/go-sql-driver/mysql" // MySQL driver
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

// Database instance
var db *sql.DB

// Connect function for MySQL
func Connect() error {
	var err error
	// Load .env file
	if err = godotenv.Load(".env"); err != nil {
		log.Println("No .env file found")
	}

	// Read environment variables
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	// Create DSN (Data Source Name)
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPassword, dbHost, dbPort, dbName)
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		return err
	}

	// Check the connection
	if err = db.Ping(); err != nil {
		return err
	}
	fmt.Println("Database connection successful!")
	return nil
}

func main() {
	// Connect to the database
	if err := Connect(); err != nil {
		log.Fatal("Database Connection Error:", err)
	}
	defer db.Close()

	// Create a Fiber instance
	app := fiber.New()

	// add logger middleware
	app.Use(logger.New())

	// Register the hello route
	routes.HelloRoute(app)

	// Start the server
	log.Fatal(app.Listen(":3000"))
}

