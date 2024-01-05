package users

import "dating-apps/domain/users/model"

type UsercaseInterface interface {
	Register(payload model.User) error
	Login(payload model.UserLogin) (string, error)
	GetProfiles(email string) (users []model.User, err error)
	Swipe(email, idTarget, action, subscription string) error
	Purchase(email string) error
}
