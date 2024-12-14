package helper

import (
	"database/sql"

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
	errFromPanic := recover() // mengambil error dari panic
	if err != nil || errFromPanic != nil {
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
