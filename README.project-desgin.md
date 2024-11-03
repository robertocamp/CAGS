# Project Design

Now that I've got the simple `hello world` code working in main.go, I see that I have routes:

`app.Get("/", hello)`

and a hello function:

```
func hello(c *fiber.Ctx) error {
	return c.SendString("Hello, World!")
}
```


The end-state project design looks like this:

```
├── README.md
├── api
│   ├── handlers
│   │   └── book_handler.go
│   ├── presenter
│   │   └── book.go
│   └── routes
│       └── book.go
├── go.mod
├── go.sum
├── main.go
└── pkg
    ├── book
    │   ├── repository.go
    │   └── service.go
    └── entities
        └── book.go
```

So before embarking on the some of the business logic API code, can we take a "baby step" and simply migrate the `"/" hello` route to the end-state design?


In this scenario, I would envisage an `api/routes/hello.go` file and then an `api/handlers/hello_handler.go` file?

Am I imagining that correctly?



When the project gets reloaded by air, I get this message:

```building...
main.go:6:2: package github/CAGS/api/routes is not in std (/opt/homebrew/Cellar/go/1.22.5/libexec/src/github/CAGS/api/routes)
```


The project initialised with `go mod init cags`

and go.mod has:

```
go 1.22.5

require (
	github.com/go-sql-driver/mysql v1.8.1
	github.com/gofiber/fiber/v2 v2.52.5
)

require (
	filippo.io/edwards25519 v1.1.0 // indirect
	github.com/andybalholm/brotli v1.0.5 // indirect
	github.com/google/uuid v1.5.0 // indirect
	github.com/joho/godotenv v1.5.1 // indirect
	github.com/klauspost/compress v1.17.0 // indirect
	github.com/mattn/go-colorable v0.1.13 // indirect
	github.com/mattn/go-isatty v0.0.20 // indirect
	github.com/mattn/go-runewidth v0.0.15 // indirect
	github.com/rivo/uniseg v0.2.0 // indirect
	github.com/valyala/bytebufferpool v1.0.0 // indirect
	github.com/valyala/fasthttp v1.51.0 // indirect
	github.com/valyala/tcplisten v1.0.0 // indirect
	golang.org/x/sys v0.15.0 // indirect
)
```



Here is main.go:

```
package main

import (
	"database/sql"
	"fmt"
	"cags/api/routes"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql" // MySQL driver
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"

	"cags/api/routes"
	"cags/api/handlers"
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

	// Register the hello route
	routes.HelloRoute(app)

	// Start the server
	log.Fatal(app.Listen(":3000"))
}
```

and here are the errors:


```
go run main.go                                             ▦ cags 🔀  project-design⎪●◦◌◦⎥ go ∩ v1.22.5  aws ▲   us-east-2 2024-11-03 15:03
# command-line-arguments
./main.go:14:2: routes redeclared in this block
        ./main.go:6:2: other declaration of routes
./main.go:14:2: "cags/api/routes" imported and not used
./main.go:15:2: "cags/api/handlers" imported and not used
```