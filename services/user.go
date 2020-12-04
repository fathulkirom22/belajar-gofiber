package services

import (
	"belajar-fiber/database"
	"belajar-fiber/model"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func GetUsers(context *fiber.Ctx) error {
	db := database.DBConn
	var users []model.User
	db.Find(&users)
	return context.JSON(users)
}

func GetUser(context *fiber.Ctx) error {
	id := context.Params("id")
	db := database.DBConn
	var user model.User
	db.Find(&user, id)
	if user.Name == "" {
		return context.Status(404).SendString(fmt.Sprintf("No User Found with ID %s", id))
	}
	return context.JSON(user)
}

func NewUser(context *fiber.Ctx) error {
	db := database.DBConn
	user := new(model.User)
	if err := context.BodyParser(user); err != nil {
		return context.Status(400).SendString(err.Error())
	}
	if err := db.Create(&user).Error; err != nil {
		return context.Status(400).SendString(err.Error())
	}
	return context.JSON(user)
}

func DeleteUser(context *fiber.Ctx) error {
	id := context.Params("id")
	db := database.DBConn

	var user model.User
	db.First(&user, id)
	if user.Name == "" {
		return context.Status(404).SendString(fmt.Sprintf("No User Found with ID %s", id))
	}
	db.Delete(&user)
	return context.SendString("User Successfully deleted")
}
