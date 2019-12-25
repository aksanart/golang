package repository

import "aksan/golangcrudmvc/models"

type BukuRepository interface {
	Save(*models.Buku) error
	Update(int, *models.Buku) error
	Delete(int) error
	FindAll() (models.Bukus, error)
	FindByID(int) (models.Buku, error)
}
