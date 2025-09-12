package handlers

import (
	"backend/config"
	"backend/models"
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
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
		)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
		}
		partners = append(partners, partner)
	}

	return c.Status(fiber.StatusOK).JSON(partners)
}