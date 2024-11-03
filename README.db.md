# MYSQL

## mysql setup:  Docker container smoke test

To set up a local MySQL Docker container for testing this code, you can use the official MySQL Docker image. Here’s a step-by-step guide:

### 1. Run a MySQL Container with Docker 

Use the following command to start a MySQL container with the necessary configurations:

```bash
docker run --name mysql_fiber_demo -p 3306:3306 -e MYSQL_ROOT_PASSWORD=password -e MYSQL_DATABASE=fiber_demo -d mysql:latest
```

This command:
- **Names the container** `mysql_fiber_demo`
- **Maps port 3306** on the container to port 3306 on your localhost (so `localhost:3306` will access MySQL)
- **Sets the root password** to `password` (`MYSQL_ROOT_PASSWORD=password`)
- **Creates a database** named `fiber_demo` (`MYSQL_DATABASE=fiber_demo`)

### 2. Update Database Configuration in `main.go`

Since MySQL typically runs on port 3306 (not 5432 as in your code), adjust the `port` constant in your code. Additionally, replace the connection string to use the correct MySQL host and port:

```go
const (
    host     = "localhost"
    port     = 3306 // MySQL default port
    user     = "root"
    password = "password"
    dbname   = "fiber_demo"
)

// Connect function
func Connect() error {
    var err error
    // Use DSN string to open
    db, err = sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", user, password, host, port, dbname))
    if err != nil {
        return err
    }
    if err = db.Ping(); err != nil {
        return err
    }
    return nil
}
```

### 3. Initialize the Database Table

The container only creates the database, not the tables. You’ll need to create the `employees` table manually before running your Fiber app.

1. Connect to the MySQL container:
   ```bash
   docker exec -it mysql_fiber_demo mysql -u root -p
   ```
   Enter the password `password` when prompted.

2. Switch to the `fiber_demo` database:
   ```sql
   USE fiber_demo;
   ```

3. Create the `employees` table:
   ```sql
   CREATE TABLE employees (
       id INT AUTO_INCREMENT PRIMARY KEY,
       name VARCHAR(255) NOT NULL,
       salary DECIMAL(10, 2) NOT NULL,
       age INT NOT NULL
   );
   ```

### 4. Run the Fiber Application

With the MySQL container running and the `employees` table created, you can now start your Fiber app:

```bash
go run main.go
```

### Testing the Endpoints

Now, you can test the API endpoints:


- **POST** a new employee:
  ```bash
  curl -X POST http://localhost:3000/employee -H "Content-Type: application/json" -d '{"name":"John Doe","salary":50000,"age":30}'
  ```

- **PUT** to update an employee:
  ```bash
  curl -X PUT http://localhost:3000/employee -H "Content-Type: application/json" -d '{"id":1,"name":"Jane Doe","salary":55000,"age":32}'
  ```

- **GET** all employees:
  ```bash
  curl http://localhost:3000/employee
  ```

- **DELETE** an employee:
  ```bash
  curl -X DELETE http://localhost:3000/employee -H "Content-Type: application/json" -d '{"id":1}'
  ```

## refactor:  convert main to msyql for web backend

