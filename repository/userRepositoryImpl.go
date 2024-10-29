package repository

import (
	"database/sql"
	"encoding/json"
	"gkru-service/entity"
	"gkru-service/helper"

	"github.com/gofiber/fiber/v2"
)

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type userRepositoryImpl struct {
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepositoryImpl{}
}

func (repository *userRepositoryImpl) FindOne(ctx *fiber.Ctx, tx *sql.Tx) (entity.User, error) {
	sqlScript := "SELECT id, username, ketua_lingkungan, ketua_wilayah FROM users WHERE username = ? AND password = ?"
	body := ctx.Body()
	request := new(LoginRequest)
	err := json.Unmarshal(body, request)
	if err != nil {
		return entity.User{}, fiber.NewError(fiber.StatusBadRequest, "Invalid request body")
	}

	result, err :=tx.Query(sqlScript, request.Username, request.Password)
	if err != nil {
		return entity.User{}, helper.CreateErrorMessage("Failed to execute query", err)
	}
	defer result.Close()
	
	user := entity.User{}
	if result.Next(){
		err := result.Scan(&user.Id, &user.Username, &user.KetuaLingkungan, &user.KetuaWilayah)
		if err != nil {
			return user, helper.CreateErrorMessage("Failed to scan result", err)
		}
		return user, nil
	} else{
		return user, fiber.NewError(fiber.StatusNotFound, "user is not found")
	}
}

func (repository *userRepositoryImpl) FindAll(ctx *fiber.Ctx, tx *sql.Tx) ([]entity.User, error) {
	sqlScript := "SELECT id, username, ketua_lingkungan, ketua_wilayah FROM users"
	
	result, err := tx.Query(sqlScript)
	if err != nil {
		return []entity.User{}, helper.CreateErrorMessage("Failed to execute query", err)
	}
	defer result.Close()

	var users []entity.User

	for result.Next() {
		user := entity.User{}
		err := result.Scan(&user.Id, &user.Username, &user.KetuaLingkungan, &user.KetuaWilayah)
		if err != nil {
			return nil, helper.CreateErrorMessage("Failed to scan result", err)
		}
		users = append(users, user)
	}

	if len(users) == 0 {
		return nil, fiber.NewError(fiber.StatusNotFound, "No users found")
	}

	return users, nil
}

