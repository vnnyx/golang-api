package helper

import (
	"database/sql"
	"golang-simple-api/exception"
)

func CommitOrRollback(tx *sql.Tx) {
	err := recover()
	if err != nil {
		errorRollback := tx.Rollback()
		exception.PanicIfNeeded(errorRollback)
		exception.PanicIfNeeded(err)
	} else {
		errorCommit := tx.Commit()
		exception.PanicIfNeeded(errorCommit)
	}
}
