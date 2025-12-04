package services

import (
	"errors"
	"fiber_rest/dal"
	"fiber_rest/types"
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
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

func GetTodo(c *fiber.Ctx) error {
	todoID := c.Params("todoID")

	d := types.TodoResponse{}
	res := dal.GetTodoByID(&d, todoID)
	if res.Error != nil {
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return c.Status(400).JSON(fiber.Map{
				"message": "Todo not found",
			})
		}
		return c.Status(500).JSON(fiber.Map{
			"message ": "Failed to get todo",
		})
	}
	return c.JSON(d)
}

func UpdateTodo(c *fiber.Ctx) error {
	todoID := c.Params("todoID")

	t := new(types.TodoUpdateDTO)
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

	res := dal.UpdateTodo(todoID, t)
	if res.Error != nil || res.RowsAffected == 0 {
		return c.Status(500).JSON(fiber.Map{
			"message": "Failed ",
		})
	}
	return c.JSON(fiber.Map{
		"message": "Todo Update Succesfully",
	})
}

func DeleteTodo(c *fiber.Ctx) error {
	todoID := c.Params("todoID")

	res := dal.DeleteTodo(todoID)
	if res.Error != nil || res.RowsAffected == 0 {
		return c.Status(500).JSON(fiber.Map{
			"message": "Failed to delete todo",
		})
	}
	return c.JSON(fiber.Map{
		"message ": "Todo Delete Succesfuly",
	})
}
