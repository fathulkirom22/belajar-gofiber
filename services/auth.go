package services

import (
	"belajar-fiber/blueprint"
	"belajar-fiber/config"
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

func getUserByEmail(email string) (*model.User, *blueprint.Error) {
	db := database.DBConn
	var user model.User
	if err := db.Where(&model.User{Email: email}).Find(&user).Error; err != nil {
		return nil, blueprint.CreateError(500, err.Error())
	}
	if user.Name == "" {
		return nil, blueprint.CreateError(404, lang.DataNotFound)
	}
	return &user, nil
}

func Login(context *fiber.Ctx) error {
	type request struct {
		Email    string `validate:"required"`
		Password string `validate:"required"`
	}

	var req request
	if err := context.BodyParser(&req); err != nil {
		return context.Status(400).SendString(err.Error())
	}
	if user, err := getUserByEmail(req.Email); err != nil {
		return context.Status(err.Code).SendString(err.Error())
	} else {
		return context.JSON(fiber.Map{
			"login": checkPasswordHash(req.Password, user.Password),
			"test":  config.Config("JWT_SECRET"),
		})
	}
}
