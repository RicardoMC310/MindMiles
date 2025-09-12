package handlers

import (
	"backend/config"
	"backend/dtos"
	"backend/models"
	"context"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

func GetAllUsers(c *fiber.Ctx) error {

	rows, err := config.Conn.Query(context.Background(),"SELECT * FROM users")
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	defer rows.Close()


	var users []models.UsersModel
	for rows.Next() {
		var u models.UsersModel
		err := rows.Scan(
			&u.ID,
			&u.Name,
			&u.Email,
			&u.Password,
			&u.Rules,
			&u.CreateAt,
		)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
		}
		users = append(users, u)
	}

	if rows.Err() != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": rows.Err().Error()})
	}

	return c.Status(fiber.StatusOK).JSON(users)
}

func CreateUser(c *fiber.Ctx) error {
	var body dtos.UsersDto

	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	if err := dtos.Validate.Struct(body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(body.Password), bcrypt.DefaultCost)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	_, err = config.Conn.Exec(
		context.Background(), 
		"INSERT INTO users(name, email, password, rules) VALUES ($1, $2, $3, $4)",
		body.Name, body.Email, hashedPassword, body.Rules)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
		}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"message": "Criado com sucesso"})
}