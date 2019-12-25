package service

import (
	"aksan/golangcrudmvc/models"
	"database/sql"
)

type bukuService struct {
	db *sql.DB
}

func NewBukuService(db *sql.DB) *bukuService {
	return &bukuService{db}
}

func (r *bukuService) Save(buku *models.Buku) error {
	query := `INSERT INTO "buku"("judul","harga","path") VALUES($1,$2,$3)`

	statement, err := r.db.Prepare(query)

	if err != nil {
		return err
	}
	defer statement.Close()
	_, err = statement.Exec(buku.Judul, buku.Harga, buku.Path)
	if err != nil {
		return err
	}
	return nil
}

func (r *bukuService) FindAll() (models.Bukus, error) {
	query := `SELECT * FROM "buku"`
	rows, err := r.db.Query(query)

	var bukus models.Bukus

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var buku models.Buku
		err = rows.Scan(&buku.Harga, &buku.Idbuku, &buku.Judul, &buku.Path)
		bukus = append(bukus, buku)
		if err != nil {
			return nil, err
		}
	}

	return bukus, nil

}

func (r *bukuService) Update(id int, buku *models.Buku) error {
	query := `UPDATE "buku" SET "judul"=$1, "harga"=$2, "path"=$3 WHERE "id_buku"=$4`

	statement, err := r.db.Prepare(query)

	if err != nil {
		return err
	}
	defer statement.Close()
	_, err = statement.Exec(buku.Judul, buku.Harga, buku.Path, id)
	if err != nil {
		return err
	}
	return nil
}

func (r *bukuService) FindByID(id int) (models.Buku, error) {
	query := `SELECT * FROM "buku" WHERE "id_buku"=$1`

	rows, err := r.db.Query(query, id)

	var buku models.Buku
	if err != nil {
		return buku, err
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&buku.Harga, &buku.Idbuku, &buku.Judul, &buku.Path)
		if err != nil {
			return buku, err
		}
	}
	return buku, nil
}

func (r *bukuService) Delete(id int) error {
	query := `DELETE FROM "buku" WHERE "id_buku"=$1`
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
