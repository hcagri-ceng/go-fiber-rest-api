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
	a := map[string]string{
		"name":    "Hasan Çağrı ",
		"surname": "Tuncer",
	}
	database.Connect()
	database.DB.AutoMigrate(dal.Todo{}) // içisine vereceğimiz structları yoksa oluşturuyor varsa güncelliyor.
	app := fiber.New()
	app.Get("/hamit", func(c *fiber.Ctx) error {
		return c.JSON(a)
	})

	// Todo App uygulaması yaptığımız için Kullanıcıdan veri almalıyız alttaki işlemler bu yüzden

	// Kullanıcı düzene aykırı veriler giremesin diye gireceği yerleri burada farklı bir struct ile tanımladık.

	// Kullanıcıdan veri alacağımız için POST metodunu kullanıyoruz.

	app.Post("/todos", services.CreateTodo)
	app.Get("/todos", services.GetTodos)
	app.Get("/todos/:todoID", services.GetTodo)
	app.Put("/todos/:todoID", services.UpdateTodo)
	app.Delete("/todos/:todoID", services.DeleteTodo)
	app.Listen("localhost:3000")
}
