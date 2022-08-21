package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/mskydream/micro/config"
	"github.com/mskydream/micro/database"
	"github.com/mskydream/micro/routes"
)

func setupRoutes(app *fiber.App) {
	api := app.Group("/api")
	api.Post("/todo", routes.CreateTodo)       // POST /api/todo
	api.Get("/todos", routes.GetTodos)         // GET /api/todo
	api.Get("/todo/:id", routes.GetTodo)       // GET /api/todo/:id
	api.Delete("/todo/:id", routes.DeleteTodo) // DELETE /api/todo/:id
	api.Put("/todo/:id", routes.UpdateTodo)    // PUT /api/todo/:id

}

func main() {
	c, err := config.LoadConfig()
	if err != nil {
		log.Fatalln(err)
	}
	database.Init(&c)

	app := fiber.New()
	setupRoutes(app)

	log.Fatal(app.Listen(c.Port))
}
