package repository

import (
	"database/sql"
	"fmt"
	"log"
)

func handleRowsAffected(res sql.Result) error {
	count, err := res.RowsAffected()
	if err == nil && count == 0 {
		err = fmt.Errorf("0 row affected")
		log.Println(err)
	}
	log.Println(err)
	return err
}
