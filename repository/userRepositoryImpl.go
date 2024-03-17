package repository

import (
	"database/sql"
	"errors"
	"gkru-service/entity"
	"gkru-service/helper"

	"github.com/gofiber/fiber/v2"
)

type userRepositoryImpl struct {
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepositoryImpl{}
}

func (repository *userRepositoryImpl) FindOne(ctx *fiber.Ctx, tx *sql.Tx) (entity.User, error) {
	sqlScript := "SELECT id, username FROM users WHERE username = ? AND password = ?"
	username := ctx.FormValue("username")
	password := ctx.FormValue("password")

	result, err :=tx.Query(sqlScript, username, password)
	helper.PanicIfError(err);
	defer result.Close()
	
	user := entity.User{}
	if result.Next(){
		err := result.Scan(&user.Id, &user.Username)
		helper.PanicIfError(err)
		return user, nil
	} else{
		return user, errors.New("User is Not Found")
	}
}
