package repository

import (
	"database/sql"
	"fmt"
	"gkru-service/entity"
	"gkru-service/helper"

	"github.com/gofiber/fiber/v2"
)

type transactionHistoryRepositoryImpl struct {
}

func NewTransactionHistoryRepository(db *sql.DB) TransactionHistoryRepository {
	return &transactionHistoryRepositoryImpl{}
}

func (repository *transactionHistoryRepositoryImpl) GetTotalIncome(ctx *fiber.Ctx, tx *sql.Tx) (entity.AmountHistory, error) {
    idWilayah := ctx.Query("id_wilayah")
    idLingkungan := ctx.Query("id_lingkungan")
    sqlScript := "SELECT SUM(nominal) FROM riwayat_transaksi WHERE keterangan = 'IN'"
	result, err := helper.AddLingkunganOrWilayahQueryHelper(idWilayah, idLingkungan, sqlScript, tx);
    helper.PanicIfError(err)
    defer result.Close()

    // Proses hasil query
    AmountHistory := entity.AmountHistory{}
    if result.Next(){
        err := result.Scan(&AmountHistory.Nominal)
        helper.PanicIfError(err)
        return AmountHistory, nil
    } else{
        return AmountHistory, fiber.NewError(fiber.StatusInternalServerError, "Error Internal")
    }
}

func (repository *transactionHistoryRepositoryImpl) GetTotalOutcome(ctx *fiber.Ctx, tx *sql.Tx) (entity.AmountHistory, error) {
	idWilayah := ctx.Query("id_wilayah")
    idLingkungan := ctx.Query("id_lingkungan")
	sqlScript := "SELECT SUM(nominal) FROM riwayat_transaksi WHERE keterangan = 'OUT'"
	result, err := helper.AddLingkunganOrWilayahQueryHelper(idWilayah, idLingkungan, sqlScript, tx);

	helper.PanicIfError(err);
	defer result.Close()
	fmt.Println("total outcome", result);
	
	AmountHistory := entity.AmountHistory{}
	if result.Next(){
		err := result.Scan(&AmountHistory.Nominal)
		helper.PanicIfError(err)
		return AmountHistory, nil
	} else{
		return AmountHistory, fiber.NewError(fiber.StatusInternalServerError , "Error Internal")
	}
}
