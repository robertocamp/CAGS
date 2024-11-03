package routes 


import (
	"cags/api/handlers"
	"github.com/gofiber/fiber/v2"
)

// helloRoute sets up the route for the "/" endpoint
func HelloRoute(app *fiber.App) {
	app.Get("/", handlers.Hello)
}	