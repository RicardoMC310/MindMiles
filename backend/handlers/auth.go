package handlers

import (
	"backend/config"
	"backend/dtos"
	"backend/models"
	"backend/utils"
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

func Login(c *fiber.Ctx) error {
	var body dtos.RequestLoginDto

	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	if err := dtos.Validate.Struct(body); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	var user models.UsersModel
	err := config.Conn.QueryRow(
		context.Background(),
		"SELECT * FROM users WHERE email = $1", body.Email).Scan(
		&user.ID, &user.Name, &user.Email, &user.Password, &user.Rules, &user.CreateAt)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password)); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid email or password"})
	}

	claims := jwt.MapClaims{
		"id":         user.ID,
		"name":       user.Name,
		"email":      user.Email,
		"rules":      user.Rules,
		"created_at": user.CreateAt,
	}

	jwtSecret := utils.GetEnv("JWT_SECRET")
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(jwtSecret))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "login sucessful",
		"token":   tokenString,
		"user": fiber.Map{
			"id":         user.ID,
			"name":       user.Name,
			"email":      user.Email,
			"rules":      user.Rules,
			"created_at": user.CreateAt,
		},
	})
}
