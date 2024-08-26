package repository

import (
	"database/sql"
	"gkru-service/entity"
	// "encoding/json"
	"gkru-service/helper"
	"github.com/gofiber/fiber/v2"
)

type dataAnggotaRepositoryImpl struct {
}

func NewDataAnggotaRepository(db *sql.DB) DataAnggotaRepository {
	return &dataAnggotaRepositoryImpl{}
}

func (repository *dataAnggotaRepositoryImpl) GetTotalAnggota(ctx *fiber.Ctx, tx *sql.Tx) (entity.TotalAnggota, error) {
	sqlScript := "SELECT COUNT(*) FROM data_anggota where status='HIDUP'"
	result, err :=tx.Query(sqlScript)
	helper.PanicIfError(err);
	defer result.Close()
	
	totalAnggota := entity.TotalAnggota{}
	if result.Next(){
		err := result.Scan(&totalAnggota.Total)
		helper.PanicIfError(err)
		return totalAnggota, nil
	} else{
		return totalAnggota, fiber.NewError(fiber.StatusInternalServerError , "Error Internal")
	}
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
