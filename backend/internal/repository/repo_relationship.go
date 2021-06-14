package repository

import (
	"app/internal/model"
	"database/sql"
	"reflect"

	"github.com/lib/pq"
)

type RelationshipRepoImpl struct {
	DB *sql.DB
}

func NewRelationshipRepo(db *sql.DB) RelationshipRepo {
	return &RelationshipRepoImpl{db}
}

func (r *RelationshipRepoImpl) selectById(id int, str string) (rels []model.Relationship, err error) {
	rows, err := r.DB.Query(str, id)
	if err != nil {
		return
	}

	e := model.Relationship{}
	s := reflect.ValueOf(&e).Elem()
	numCols := s.NumField()
	columns := make([]interface{}, numCols)
	for i := 0; i < numCols; i++ {
		field := s.Field(i)
		columns[i] = field.Addr().Interface()
	}

	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(columns...)
		if err != nil {
			panic(err)
		}
		rels = append(rels, e)
	}

	return rels, nil
}

func (r *RelationshipRepoImpl) Select(id int) (rels []model.Relationship, err error) {
	rels, err = r.selectById(id, "select * from relationship where user1=$1")
	return
}

func (r *RelationshipRepoImpl) Friends(id int) (rels []model.Relationship, err error) {
	rels, err = r.selectById(id, "select * from relationship where user1=$1 and type='friend'")
	return
}

func (r *RelationshipRepoImpl) Requests(id int) (rels []model.Relationship, err error) {
	rels, err = r.selectById(id, "select * from relationship where user1=$1 and type='request'")
	return
}

func (r *RelationshipRepoImpl) FriendsDetail(id int) (fd string, err error) {
	err = r.DB.QueryRow("select friends_json($1)", id).Scan(&fd)
	return
}

func (r *RelationshipRepoImpl) MutualFriends(u1, u2 int) (mf []int64, err error) {
	row := r.DB.QueryRow("select mutual_friends($1, $2)", u1, u2)
	var arr pq.Int64Array
	err = row.Scan(&arr)
	mf = arr
	return
}

func (r *RelationshipRepoImpl) ChangeType(u1, u2 int, t string) (err error) {
	query := `insert into relationship(user1, user2, type) values($1, $2, $3)
	on conflict (user1, user2) do update set type=$3`
	res, err := r.DB.Exec(query, u1, u2, t)
	if err == nil {
		err = handleRowsAffected(res)
	}
	return
}

func (r *RelationshipRepoImpl) SelectRelationshipWith(u1, u2 int) (t string) {
	err := r.DB.QueryRow("select type from relationship where user1=$1 and user2=$2", u1, u2).Scan(&t)
	if err != nil {
		return ""
	}
	return
}

func (r *RelationshipRepoImpl) Delete(u1, u2 int) (err error) {
	query := `delete from relationship where user1=$1 and user2=$2`
	res, err := r.DB.Exec(query, u1, u2)
	if err == nil {
		err = handleRowsAffected(res)
	}
	return
}
