Yes, Go Fiber supports adding middleware for enhanced logging, including recording API calls, response times, and status codes. You can use the built-in `logger` middleware to achieve this. The `logger` middleware logs each request to the terminal with information like HTTP method, route path, status code, and response time.

Here’s how to set it up for your project:

### Step 1: Add the `Logger` Middleware

In `main.go`, import Fiber’s logger middleware and add it to your application setup:

```go
package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"cags/api/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
	_ "github.com/go-sql-driver/mysql"
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

	// Add Logger middleware
	app.Use(logger.New())

	// Register the hello route
	routes.HelloRoute(app)

	// Start the server
	log.Fatal(app.Listen(":3000"))
}
```

### Step 2: Customize Logger (Optional)

Fiber’s `logger` middleware supports customization to control the logging output format. For example:

```go
app.Use(logger.New(logger.Config{
	Format:     "[${time}] ${status} - ${method} ${path}\n",
	TimeFormat: "2006-01-02 15:04:05",
	TimeZone:   "Local",
}))
```

In this configuration:
- **`${time}`**: Displays the current time of the request.
- **`${status}`**: Shows the HTTP status code.
- **`${method}`**: Logs the HTTP method (GET, POST, etc.).
- **`${path}`**: Shows the requested path.

You can further customize the format as needed, adding elements like `${latency}` (to log response time), `${ip}` (client IP address), and more.

### Result

With this setup, each API request will be logged to the terminal in real time, including useful details like HTTP method, status code, and path. For example, the output might look like:

```
[2024-11-03 15:15:12] 200 - GET /hello
[2024-11-03 15:15:13] 404 - GET /unknown
```

This provides clear visibility into API activity, which is helpful for development and debugging. Let me know if you’d like to explore more advanced logging features or formats!