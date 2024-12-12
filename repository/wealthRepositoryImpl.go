package repository

import (
	"database/sql"
	"gkru-service/entity"
	"gkru-service/helper"

	"github.com/gofiber/fiber/v2"
)

type wealthRepositoryImpl struct {
}

func NewWealthRepository(db *sql.DB) WealthRepository {
	return &wealthRepositoryImpl{}
}

// kemungkinan ini gak dipake, pake aja count langsung di th history

func (repository *wealthRepositoryImpl) GetTotal(ctx *fiber.Ctx, tx *sql.Tx) (entity.TotalWealth, error) {
	sqlScript := "SELECT SUM(total) FROM wealth"
	result, err := tx.Query(sqlScript)
	helper.PanicIfError(err)
	defer result.Close()

	totalWealth := entity.TotalWealth{}
	if result.Next() {
		err := result.Scan(&totalWealth.Total)
		if err != nil {
			return totalWealth, helper.CreateErrorMessage("Gagal untuk scan result", err)
		}
		return totalWealth, nil
	} else {
		return totalWealth, fiber.NewError(fiber.StatusInternalServerError, "Error Internal")
	}
}
