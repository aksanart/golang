package repository

import "aksan/golangcrudmvc/models"

type RestRepository interface {
	Save(*models.UserRest) error
	Update(int, *models.UserRest) error
	Delete(int) error
	FindAll() (models.UserRests, error)
	FindByID(int) (models.UserRest, error)
	FindByUsernameRest(string) (models.UserRest, error)
}
