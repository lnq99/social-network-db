package repository

import (
	"app/pkg/logger"
	"database/sql"
	"fmt"
)

func handleRowsAffected(res sql.Result) error {
	count, err := res.RowsAffected()
	if err == nil && count == 0 {
		err = fmt.Errorf("0 row affected")
		logger.Info(err)
	}
	logger.Info(err)
	return err
}
