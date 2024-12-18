package helper

import (
	"database/sql"
	"fmt"

	"github.com/sirupsen/logrus"
)

func CommitOrRollback(tx *sql.Tx, logger *logrus.Logger) {
	err := recover()
	if err != nil {
		errorRollback := tx.Rollback()
		if errorRollback != nil {
			logger.WithError(errorRollback).Warn("Error when rolling back transaction")
		}
	} else {
		errorCommit := tx.Commit()
		if errorCommit != nil {
			logger.WithError(errorCommit).Warn("Error when commiting transaction")
		}
	}
}

//buat mengatasi rollback transaction jika ada error ditenbgah-tengah tapi tidak memanggil panic
// kerena recover itu mengecek error dari panic makanya dia selalu ter commit jika ada error non panic gitu
func CommitOrRollback2(tx *sql.Tx, logger *logrus.Logger, err error) {
	fmt.Println("masuk sini")
    defer func() {
        if r := recover(); r != nil {
			fmt.Println("masuk sini 1")
            logger.WithField("panic", r).Error("Recovered from panic in transaction")
            _ = tx.Rollback() // Rollback jika terjadi panic
        }
    }()
    if err != nil {
		fmt.Println("masuk sini 2")
        logger.WithError(err).Warn("Rolling back transaction due to error")
        rollbackErr := tx.Rollback()
        if rollbackErr != nil {
            logger.WithError(rollbackErr).Error("Failed to rollback transaction")
        }
    } else {
		fmt.Println("masuk sini 3")
        commitErr := tx.Commit()
        if commitErr != nil {
            logger.WithError(commitErr).Warn("Error committing transaction")
        }
    }
}
