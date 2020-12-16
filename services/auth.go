package services

import (
	"belajar-fiber/blueprint"
	"belajar-fiber/database"
	"belajar-fiber/lang"
	"belajar-fiber/model"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

func checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func getUserByEmail(user *model.User, email string) *blueprint.Error {
	db := database.DBConn

	if err := db.Where(&model.User{Email: email}).Find(&user).Error; err != nil {
		return blueprint.CreateError(500, err.Error())
	}

	if user.Name == "" {
		return blueprint.CreateError(404, lang.DataNotFound)
	}

	return nil
}

// Login ...
func Login(context *fiber.Ctx) error {
	type request struct {
		Email    string `validate:"required"`
		Password string `validate:"required"`
	}

	var req request
	if err := context.BodyParser(&req); err != nil {
		return context.Status(400).SendString(err.Error())
	}

	var user model.User
	if err := getUserByEmail(&user, req.Email); err != nil {
		return context.Status(err.Code).SendString(err.Error())
	}

	if !checkPasswordHash(req.Password, user.Password) {
		return context.Status(401).SendString("Invalid password")
	}

	return context.SendString("Success login")
}
