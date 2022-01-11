package repository

import (
	"database/sql"
	"fmt"

	"app/pkg/logger"
)

func handleRowsAffected(res sql.Result) error {
	count, err := res.RowsAffected()
	if err == nil && count == 0 {
		err = fmt.Errorf("0 row affected")
		logger.Info(err)
	}
	return err
}
