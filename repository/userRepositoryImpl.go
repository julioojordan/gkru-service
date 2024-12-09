package repository

import (
	"database/sql"
	"encoding/json"
	"gkru-service/entity"
	"gkru-service/helper"
	"strconv"
	"strings"
	"time"

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

	result, err := tx.Query(sqlScript, request.Username, request.Password)
	if err != nil {
		return entity.User{}, helper.CreateErrorMessage("Failed to execute query", err)
	}
	defer result.Close()

	user := entity.User{}
	if result.Next() {
		err := result.Scan(&user.Id, &user.Username, &user.KetuaLingkungan, &user.KetuaWilayah)
		if err != nil {
			return user, helper.CreateErrorMessage("Failed to scan result", err)
		}
		return user, nil
	} else {
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

	// if len(users) == 0 {
	// 	return nil, fiber.NewError(fiber.StatusNotFound, "No users found")
	// }

	return users, nil
}

func (repository *userRepositoryImpl) Update(ctx *fiber.Ctx, tx *sql.Tx) (entity.User, error) {
	sqlScript := "UPDATE users SET"
	idUser, err := strconv.Atoi(ctx.Params("idUser"))
	if err != nil {
		return entity.User{}, fiber.NewError(fiber.StatusBadRequest, "Invalid id user, it must be an integer")
	}
	body := ctx.Body()
	request := new(helper.UserRequest)
	marshalError := json.Unmarshal(body, request)
	if marshalError != nil {
		return entity.User{}, fiber.NewError(fiber.StatusBadRequest, "Invalid request body")
	}

	currentTime := time.Now()
	var params []interface{}
	var setClauses []string

	if request.Username != "" {
		setClauses = append(setClauses, "username = ?")
		params = append(params, request.Username)
	}
	if request.Password != "" {
		setClauses = append(setClauses, "password = ?")
		params = append(params, request.Password)
	}
	setClauses = append(setClauses, "ketua_lingkungan = ?")
	params = append(params, request.KetuaLingkungan)
	setClauses = append(setClauses, "ketua_wilayah = ?")
	params = append(params, request.KetuaWilayah)
	setClauses = append(setClauses, "updated_by = ?")
	params = append(params, request.UpdatedBy)
	setClauses = append(setClauses, "updated_date = ?")
	params = append(params, currentTime)

	// Check if there are fields to update
	if len(setClauses) == 0 {
		return entity.User{}, fiber.NewError(fiber.StatusBadRequest, "Error No fields to update")
	}

	// Joining all set clauses
	sqlScript += " " + strings.Join(setClauses, ", ") + " WHERE id = ?"
	params = append(params, idUser)

	// Executing the update statement
	_, err = tx.Exec(sqlScript, params...)
	if err != nil {
		return entity.User{}, helper.CreateErrorMessage("Failed to update data user", err)
	}

	response := entity.User{
		Id:              int32(idUser),
		Username:        request.Username,
		KetuaLingkungan: request.KetuaLingkungan,
		KetuaWilayah:    request.KetuaWilayah,
	}

	return response, nil
}

func (repository *userRepositoryImpl) Add(ctx *fiber.Ctx, tx *sql.Tx) (entity.IdInt, error) {
	sqlScript := "INSERT INTO users(username, password, ketua_lingkungan, ketua_wilayah, created_date, updated_date, created_by, updated_by) VALUES(?, ?, ?, ?, ?, ?, ?, ?)"
	body := ctx.Body()
	request := new(helper.AddUserRequest)
	err := json.Unmarshal(body, request)
	if err != nil {
		return entity.IdInt{}, fiber.NewError(fiber.StatusBadRequest, "Invalid request body")
	}

	currentTime := time.Now()

	result, err := tx.Exec(sqlScript, request.Username, request.Password, request.KetuaLingkungan, request.KetuaWilayah, currentTime, currentTime, request.CreatedBy, request.UpdatedBy)
	if err != nil {
		return entity.IdInt{}, helper.CreateErrorMessage("Failed to insert data user", err)
	}

	lastInsertId, err := result.LastInsertId()
	if err != nil {
		return entity.IdInt{}, helper.CreateErrorMessage("Failed to retrieve last inserted ID", err)
	}

	response := entity.IdInt{
		Id:             int32(lastInsertId),
	}

	return response, nil
}
