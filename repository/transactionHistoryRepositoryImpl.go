package repository

import (
	"database/sql"
	"fmt"
	"gkru-service/entity"
	"gkru-service/helper"

	"github.com/gofiber/fiber/v2"
)

type THRequest struct {
	IdWilayah string `json:"id_wilayah"`
	IdLingkungan string `json:"id_lingkungan"`
}

type transactionHistoryRepositoryImpl struct {
}

func NewTransactionHistoryRepository(db *sql.DB) TransactionHistoryRepository {
	return &transactionHistoryRepositoryImpl{}
}

func (repository *transactionHistoryRepositoryImpl) GetTotalIncome(ctx *fiber.Ctx, tx *sql.Tx) (entity.AmountHistory, error) {
	sqlScript := "SELECT SUM(nominal) FROM riwayat_transaksi WHERE keterangan = 'IN'"
	result, err :=tx.Query(sqlScript)
	helper.PanicIfError(err);
	defer result.Close()
	fmt.Println("total income", result);
	
	AmountHistory := entity.AmountHistory{}
	if result.Next(){
		err := result.Scan(&AmountHistory.Nominal)
		helper.PanicIfError(err)
		return AmountHistory, nil
	} else{
		return AmountHistory, fiber.NewError(fiber.StatusInternalServerError , "Error Internal")
	}
}

func (repository *transactionHistoryRepositoryImpl) GetTotalOutcome(ctx *fiber.Ctx, tx *sql.Tx) (entity.AmountHistory, error) {
	sqlScript := "SELECT SUM(nominal) FROM riwayat_transaksi WHERE keterangan = 'OUT'"
	result, err :=tx.Query(sqlScript)
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
