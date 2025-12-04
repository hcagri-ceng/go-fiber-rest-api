package services

import (
	"fiber_rest/dal"
	"fiber_rest/types"
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

var validate = validator.New()

func CreateTodo(c *fiber.Ctx) error {
	// t diye bir değişkeni standarta uygun bir şekilde oluşturduk.
	t := new(types.TodoCreateDTO)
	// kullanıcı farklı veri giremesin diye parselledik.
	err := c.BodyParser(t)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message ": "Hatalı Giriş",
		})
	}

	if err := validate.Struct(t); err != nil {
		valErr := err.(validator.ValidationErrors)[0]
		message := fmt.Sprintf("Field : '%s' , failed on '%s' with your value : '%s '", valErr.Field(), valErr.Tag(), valErr.Value())
		return c.Status(400).JSON(fiber.Map{
			"message ": message,
		})
	}

	// yukarıda verilerimizi aldık ama yukarıda aldığımız veriler  bizim tabloya kaydedeceğimiz şey değildi. Sadece alış formatımızdı. Doğru formatta aldığımız verileri
	// dal packagendaki asıl kaydedilecek structa atıyoruz. ID ve TAMAMLANMA zaten otomatik olduğundan sadece title'ı yolluyoruz. newTodo ise kaydedilecek değişken.
	newTodo := dal.Todo{
		Title: t.Title,
	}

	res := dal.CreateTodo(&newTodo)
	if res.Error != nil {
		return c.Status(500).JSON(fiber.Map{
			"message ": "Veritabanında problem yaşandı",
		})
	}
	return c.JSON(fiber.Map{
		"Message": "Başarılı",
	})

}

func GetTodos(c *fiber.Ctx) error {
	todos := []types.TodoResponse{}
	res := dal.GetTodos(&todos)
	if res.Error != nil {
		return c.Status(500).JSON(fiber.Map{
			"message ": "Veritabanında problem yaşandı",
		})
	}
	return c.JSON(todos)
}
