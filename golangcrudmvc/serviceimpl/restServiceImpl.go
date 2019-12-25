package serviceimpl

import (
	"aksan/golangcrudmvc/models"
	"aksan/golangcrudmvc/repository"
)

func SaveRest(b *models.UserRest, repo repository.RestRepository) error {
	err := repo.Save(b)
	if err != nil {
		return err
	}
	return nil
}

func FindAllUserRest(repo repository.RestRepository) (models.UserRests, error) {
	UserRests, err := repo.FindAll()
	if err != nil {
		return nil, err
	}
	return UserRests, err
}

func UpdateUserRest(id int, b *models.UserRest, repo repository.RestRepository) error {
	err := repo.Update(id, b)
	if err != nil {
		return err
	}
	return nil
}

func FindByIDUserRest(id int, repo repository.RestRepository) (models.UserRest, error) {
	UserRest, err := repo.FindByID(id)
	if err != nil {
		return UserRest, err
	}
	return UserRest, nil
}

func FindByUsernameRestImpl(username string, repo repository.RestRepository) (models.UserRest, error) {
	UserRest, err := repo.FindByUsernameRest(username)
	if err != nil {
		return UserRest, err
	}
	return UserRest, nil
}

func DeleteUserRest(id int, repo repository.RestRepository) error {
	err := repo.Delete(id)
	if err != nil {
		return err
	}
	return nil
}
