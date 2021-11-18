package main

import (
	"github.com/fanama/go-svelte/Api/router"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	app := fiber.New()

	app.Use(cors.New())

	app.Get("/database/:database", router.GetData)
	app.Post("/database/:database", router.InsertData)
	app.Delete("/database/:database", router.DeleteData)

	app.Listen(":8090")
}
