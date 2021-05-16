package repository

import (
	"app/internal/model"
	"database/sql"
	"reflect"
)

type ProfileRepoImpl struct {
	DB *sql.DB
}

func NewProfileRepo(db *sql.DB) ProfileRepo {
	return &ProfileRepoImpl{db}
}

func (u *ProfileRepoImpl) Select(id int) (profile model.Profile, err error) {
	row := u.DB.QueryRow("select * from Profile where id=$1 limit 1", id)

	s := reflect.ValueOf(&profile).Elem()
	numCols := s.NumField()
	columns := make([]interface{}, numCols)
	for i := 0; i < numCols; i++ {
		field := s.Field(i)
		columns[i] = field.Addr().Interface()
	}
	err = row.Scan(columns...)

	return
}

// func (u *ProfileRepoImpl) Insert(user model.Profile) error {
// 	insertStatement := `
// 	INSERT INTO users (id, name, gender, email)
// 	VALUES ($1, $2, $3, $4)`

// 	_, err := u.DB.Exec(insertStatement, user.Id, user.Name, user.Describle)
// 	if err != nil {
// 		return err
// 	}

// 	fmt.Println("Record added: ", user)

// 	return nil
// }
