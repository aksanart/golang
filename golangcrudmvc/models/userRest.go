package models

type UserRest struct {
	IdUser   int
	Username string
	Key      string
	Hash     string
}
type UserRests []UserRest

func SetUserRest() *UserRest {
	return &UserRest{}
}
