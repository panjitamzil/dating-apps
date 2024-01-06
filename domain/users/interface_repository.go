package users

import "dating-apps/domain/users/model"

type RepoInterface interface {
	Insert(payload model.User) error
	SelectUsers(email string) (users []model.User, err error)
	SelectByCondition(params *model.User) (users []model.User, err error)
	UpdateProfile(req model.User) error

	SetKey(key, value string) error
	Exist(key string) (int64, error)
	GetKeys(key, email string) (int, error)
}
