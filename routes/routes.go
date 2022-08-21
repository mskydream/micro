package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mskydream/micro/database"
	"github.com/mskydream/micro/models"
)

type Response struct {
	IsSuccess bool        `json:"isSuccess"`
	Message   string      `json:"message"`
	Data      interface{} `json:"data"`
}

func CreateTodo(c *fiber.Ctx) error {
	var todo models.Todo

	if err := c.BodyParser(&todo); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(Response{IsSuccess: false, Message: err.Error()})
	}
	database.Database.Db.Create(&todo)
	return c.Status(fiber.StatusCreated).JSON(Response{IsSuccess: true, Message: "запись успешно создана", Data: todo})
}

func GetTodos(c *fiber.Ctx) error {
	var todos []models.Todo
	database.Database.Db.Find(&todos)
	return c.Status(fiber.StatusOK).JSON(Response{IsSuccess: true, Message: "записи успешно получены", Data: todos})
}

func GetTodo(c *fiber.Ctx) error {
	var todo models.Todo
	id := c.Params("id")
	database.Database.Db.First(&todo, id)
	return c.Status(fiber.StatusOK).JSON(Response{IsSuccess: true, Message: "запись успешно получена", Data: todo})
}

func DeleteTodo(c *fiber.Ctx) error {
	var todo models.Todo
	id := c.Params("id")
	database.Database.Db.First(&todo, id)
	database.Database.Db.Delete(&todo)
	return c.Status(fiber.StatusOK).JSON(Response{IsSuccess: true, Message: "запись успешно удалена"})
}

func UpdateTodo(c *fiber.Ctx) error {
	var todo models.Todo
	id := c.Params("id")
	database.Database.Db.First(&todo, id)
	if err := c.BodyParser(&todo); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(Response{IsSuccess: false, Message: err.Error()})
	}
	database.Database.Db.Save(&todo)
	return c.Status(fiber.StatusOK).JSON(Response{IsSuccess: true, Message: "запись успешно обновлена", Data: todo})
}
