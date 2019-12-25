package serviceimpl

import (
	"aksan/golangcrudmvc/models"
	"aksan/golangcrudmvc/repository"
)

func SaveBuku(b *models.Buku, repo repository.BukuRepository) error {
	err := repo.Save(b)
	if err != nil {
		return err
	}
	return nil
}

func FindAllBuku(repo repository.BukuRepository) (models.Bukus, error) {
	bukus, err := repo.FindAll()
	if err != nil {
		return nil, err
	}
	return bukus, err
}

func Update(id int, b *models.Buku, repo repository.BukuRepository) error {
	err := repo.Update(id, b)
	if err != nil {
		return err
	}
	return nil
}

func FindByID(id int, repo repository.BukuRepository) (models.Buku, error) {
	buku, err := repo.FindByID(id)
	if err != nil {
		return buku, err
	}
	return buku, nil
}

func Delete(id int, repo repository.BukuRepository) error {
	err := repo.Delete(id)
	if err != nil {
		return err
	}
	return nil
}
