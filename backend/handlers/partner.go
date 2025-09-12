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

func GetAllPartners(c *fiber.Ctx) error {
	user := c.Locals("user").(jwt.MapClaims)
	if rules, ok := user["rules"].(string); ok && rules != "admin"{
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "acesso negado"})
	}

	rows, err := config.Conn.Query(context.Background(),"SELECT * FROM partners")
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	defer rows.Close()

	var partners []models.PartnerModel
	for rows.Next() {
		var partner models.PartnerModel
		err := rows.Scan(
			&partner.ID,
			&partner.Name,
			&partner.Email,
			&partner.CPF,
			&partner.Tel,
			&partner.Age,
			&partner.Password,
		)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
		}
		partners = append(partners, partner)
	}

	return c.Status(fiber.StatusOK).JSON(partners)
}

func RegisterNewPartner(c *fiber.Ctx) error {
	user := c.Locals("user").(jwt.MapClaims)
	if rules, ok := user["rules"].(string); ok && rules != "admin" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "admin not found!"})
	}

	var body dtos.PartnerDto

	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	if err := utils.Validate.Struct(body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(body.Password), bcrypt.DefaultCost)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	_, err = config.Conn.Exec(context.Background(), 
				"INSERT INTO partners(name, email, password, cpf, age, tel) VALUES($1,$2,$3,$4,$5,$6)",
			body.Name, body.Email, hashedPassword, body.CPF, body.Age, body.Tel)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"message": "Partner register of sucess"})
}