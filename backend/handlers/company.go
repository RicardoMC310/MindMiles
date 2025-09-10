package handlers

import (
	"backend/config"
	"backend/models"
	"context"

	"github.com/gofiber/fiber/v2"
)

func GetAllCompanys(c *fiber.Ctx) error { // precisa ser *fiber.Ctx
	rows, err := config.Conn.Query(context.Background(),
		`SELECT * 
		FROM company`,
	)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	defer rows.Close()

	var companies []models.CompanyModel
	for rows.Next() {
		var company models.CompanyModel
		err := rows.Scan(
			&company.ID,
			&company.Name,
			&company.RazaoSocial,
			&company.CNPJ,
			&company.NaturezaJuridica,
			&company.IE,
			&company.IM,
			&company.Email,
			&company.Tel,
			&company.Site,
			&company.CreatedAt,
			&company.UpdatedAt,
		)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}
		companies = append(companies, company)
	}

	return c.Status(fiber.StatusOK).JSON(companies)
}
