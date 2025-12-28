package routes

import (
	"laporan-keuangan/controllers"
	"laporan-keuangan/middleware"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	auth := app.Group("/auth")
	auth.Post("/register", controllers.Register)
	auth.Post("/login", controllers.Login)

	api := app.Group("/api", middleware.Protected())
	api.Get("/book", controllers.GetBooks)
	api.Post("/book", controllers.CreateBook)
	api.Get("/my/book", controllers.GetMyBook)

	api.Get("/budget", controllers.GetBudgetReport)
	api.Post("/budget", controllers.AddBudgetEntry)
	api.Delete("/budget/:id", controllers.DeleteBudgetEntry)
}
