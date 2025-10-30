package main

import (
	"fmt"
	"log"
	"os"

	"minify/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
)

func setupRoutes(app *fiber.App) {
	// Health check endpoint
	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"status":  "ok",
			"service": "Minify URL Shortener API",
			"version": "1.0.0",
		})
	})

	app.Get("/:url", routes.ResolveURL)
	app.Post("/api/v1", routes.ShortenURL)
}
func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file", err)
	}
	app := fiber.New()

	// Enable CORS for frontend communication
	app.Use(cors.New(cors.Config{
		AllowOrigins: "https://minify-links.vercel.app, http://localhost:5173, http://localhost:3000",
		AllowHeaders: "Origin, Content-Type, Accept",
		AllowMethods: "GET, POST, PUT, DELETE, OPTIONS",
	}))

	// Explicitly handle preflight OPTIONS requests to ensure browsers receive
	// the required Access-Control-* headers even if other handlers return errors.
	app.Options("/*", func(c *fiber.Ctx) error {
		c.Set("Access-Control-Allow-Origin", "https://minify-links.vercel.app, http://localhost:5173, http://localhost:3000")
		c.Set("Access-Control-Allow-Headers", "Origin, Content-Type, Accept")
		c.Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		return c.SendStatus(fiber.StatusNoContent)
	})

	app.Use(logger.New())
	setupRoutes(app)

	// Use APP_PORT env var if provided (e.g. ":3000"). If not set, fall back
	// to a sensible default so local runs and some hosting environments work.
	port := os.Getenv("APP_PORT")
	if port == "" {
		port = ":3000"
	}
	log.Fatal(app.Listen(port))
}
