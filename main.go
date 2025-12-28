package main

import (
	"laporan-keuangan/config"
	"laporan-keuangan/routes"
	"laporan-keuangan/utils"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	config.Connect()
	app := fiber.New()
	app.Use(cors.New())
	app.Use(logger.New())
	routes.SetupRoutes(app)
	app.Use(func(c *fiber.Ctx) error {
		return utils.ResponseError(c, 404, "endpoint tidak ditemukan")
	})
	port := os.Getenv("PORT")
	if port == "" {
		port = "6100"
	}
	log.Fatal(app.Listen(":" + port))
}
