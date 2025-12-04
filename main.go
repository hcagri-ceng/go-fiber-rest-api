package main

import (
	"fiber_rest/dal"
	"fiber_rest/database"
	"fiber_rest/services"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

var validate = validator.New()

func main() {

	database.Connect()
	database.DB.AutoMigrate(dal.Todo{}) //
	app := fiber.New()
	app.Post("/todos", services.CreateTodo)
	app.Get("/todos", services.GetTodos)
	app.Get("/todos/:todoID", services.GetTodo)
	app.Put("/todos/:todoID", services.UpdateTodo)
	app.Delete("/todos/:todoID", services.DeleteTodo)
	app.Listen("localhost:3000")
}
