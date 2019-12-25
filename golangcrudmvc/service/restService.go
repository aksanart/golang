package service

import (
	"aksan/golangcrudmvc/models"
	"database/sql"
)

type restService struct {
	db *sql.DB
}

func NewRestService(db *sql.DB) *restService {
	return &restService{db}
}

func (r *restService) Save(userrest *models.UserRest) error {
	query := `INSERT INTO "user_rest"("username","key","hash") VALUES($1,$2,$3)`

	statement, err := r.db.Prepare(query)

	if err != nil {
		return err
	}
	defer statement.Close()
	_, err = statement.Exec(userrest.Username, userrest.Key, userrest.Hash)
	if err != nil {
		return err
	}
	return nil
}

func (r *restService) FindAll() (models.UserRests, error) {
	query := `SELECT * FROM "user_rest"`
	rows, err := r.db.Query(query)

	var userRests models.UserRests

	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var userRest models.UserRest
		err = rows.Scan(&userRest.IdUser, &userRest.Username, &userRest.Key, &userRest.Hash)
		userRests = append(userRests, userRest)
		if err != nil {
			return nil, err
		}
	}

	return userRests, nil

}

func (r *restService) Update(id int, UserRest *models.UserRest) error {
	query := `UPDATE "user_rest" SET "username"=$1, "key"=$2, "hash"=$3 WHERE "id_user"=$4`

	statement, err := r.db.Prepare(query)

	if err != nil {
		return err
	}
	defer statement.Close()
	_, err = statement.Exec(UserRest.Username, UserRest.Key, UserRest.Hash, id)
	if err != nil {
		return err
	}
	return nil
}

func (r *restService) FindByID(id int) (models.UserRest, error) {
	query := `SELECT * FROM "user_rest" WHERE "id_user"=$1`

	rows, err := r.db.Query(query, id)

	var userRest models.UserRest
	if err != nil {
		return userRest, err
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&userRest.IdUser, &userRest.Username, &userRest.Key, &userRest.Hash)
		if err != nil {
			return userRest, err
		}
	}
	return userRest, nil
}

func (r *restService) Delete(id int) error {
	query := `DELETE FROM "user_rest" WHERE "id_user"=$1`
	statement, err := r.db.Prepare(query)

	if err != nil {
		return err
	}
	defer statement.Close()
	_, err = statement.Exec(id)
	if err != nil {
		return err
	}
	return nil
}

func (r *restService) FindByUsernameRest(username string) (models.UserRest, error) {
	query := `SELECT * FROM "user_rest" WHERE "username"=$1`

	rows, err := r.db.Query(query, username)

	var userRest models.UserRest
	if err != nil {
		return userRest, err
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&userRest.IdUser, &userRest.Username, &userRest.Key, &userRest.Hash)
		if err != nil {
			return userRest, err
		}
	}
	return userRest, nil
}
