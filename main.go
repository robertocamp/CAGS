package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	_ "github.com/go-sql-driver/mysql" // MySQL driver
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

	// fiber instance and HelloWorld
	app := fiber.New()


	// routes
	app.Get("/", hello)

	// start server
	log.Fatal(app.Listen(":9000"))
}

// handler function
func hello(c *fiber.Ctx) error {
	return c.SendString("Hello, World!")
}

