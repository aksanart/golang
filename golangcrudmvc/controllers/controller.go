package controllers

import (
	"aksan/golangcrudmvc/config"
	"aksan/golangcrudmvc/models"
	"aksan/golangcrudmvc/service"
	"aksan/golangcrudmvc/serviceimpl"
)

func Proses(buku *models.Buku) error {
	db, err := config.GetPostgresDB()
	if err != nil {
		return (err)
	}
	bukuService := service.NewBukuService(db)

	err = serviceimpl.SaveBuku(buku, bukuService)

	if err != nil {
		return (err)
	}
	return nil
}

func List() (models.Bukus, error) {
	db, err := config.GetPostgresDB()
	if err != nil {
		return nil, err
	}
	bukuService := service.NewBukuService(db)
	bukus, err := serviceimpl.FindAllBuku(bukuService)

	if err != nil {
		return nil, err
	}
	return bukus, nil
}

func Edit(id int, buku *models.Buku) error {
	db, err := config.GetPostgresDB()

	bukuService := service.NewBukuService(db)
	err = serviceimpl.Update(id, buku, bukuService)
	if err != nil {
		return err
	}
	return nil
}

func FindByID(id int) (models.Buku, error) {
	db, err := config.GetPostgresDB()
	var buku models.Buku
	if err != nil {
		return buku, err
	}
	bukuService := service.NewBukuService(db)
	buku, err = serviceimpl.FindByID(id, bukuService)

	return buku, err
}

func Hapus(id int) error {
	db, err := config.GetPostgresDB()
	if err != nil {
		return err
	}
	bukuService := service.NewBukuService(db)
	err = serviceimpl.Delete(id, bukuService)
	if err != nil {
		return err
	}
	return nil
}
