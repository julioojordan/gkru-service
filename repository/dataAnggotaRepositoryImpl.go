package repository

import (
	"database/sql"
	// "encoding/json"
	// "gkru-service/entity"
	// "gkru-service/helper"

	// "github.com/gofiber/fiber/v2"
)

type dataAnggotaRepositoryImpl struct {
}

func NewDataAnggotaRepository(db *sql.DB) DataAnggotaRepository {
	return &dataAnggotaRepositoryImpl{}
}

// func (repository *dataAnggotaRepositoryImpl) FindKeluargaAnggotaRel([]ids int32, tx *sql.Tx) ([]entity.DataAnggota, error) {
// 	sqlScript := "SELECT id, username FROM users WHERE username = ? AND password = ?"

// 	result, err :=tx.Query(sqlScript, request.Username, request.Password)
// 	helper.PanicIfError(err);
// 	defer result.Close()
	
// 	user := entity.User{}
// 	if result.Next(){
// 		err := result.Scan(&user.Id, &user.Username)
// 		helper.PanicIfError(err)
// 		return user, nil
// 	} else{
// 		return user, fiber.NewError(fiber.StatusNotFound, "user is not found")
// 	}
// }
