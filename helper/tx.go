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
