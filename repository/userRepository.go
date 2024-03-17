package repository

//interface repository

import (
	"database/sql"
	"gkru-service/entity"

	"github.com/gofiber/fiber/v2"
)

type UserRepository interface {
	FindOne(ctx *fiber.Ctx, tx *sql.Tx) (entity.User, error)
}