package usecase

import (
	"database/sql"
	"dating-apps/domain/users"
	"dating-apps/domain/users/model"
	"dating-apps/helpers"
	"errors"
	"fmt"
)

type Service struct {
	UserRepo users.RepoInterface
}

func NewService(UserRepo users.RepoInterface) users.UsercaseInterface {
	return &Service{
		UserRepo: UserRepo,
	}
}

func (s *Service) Register(payload model.User) error {
	var valid = false

	// Validate email
	_, err := s.UserRepo.SelectByCondition(&model.User{
		Email: payload.Email,
	})
	if err != nil {
		if err == sql.ErrNoRows {
			valid = true
		} else {
			return errors.New(helpers.ERR_VALIDATE_EMAIL + ":" + err.Error())
		}
	}

	// Encrypt password
	payload.Password = helpers.HashPass(payload.Password)

	if valid {
		// create a new user
		err = s.UserRepo.Insert(payload)
		if err != nil {
			return errors.New(helpers.ERR_CREATE_USER + ":" + err.Error())
		}

		return nil
	}

	return errors.New(helpers.ERR_EXIST)
}

func (s *Service) Login(payload model.UserLogin) (string, error) {
	// Get users info
	users, err := s.UserRepo.SelectByCondition(&model.User{
		Email: payload.Email,
	})
	if err != nil {
		if err == sql.ErrNoRows {
			return helpers.EMPTY, errors.New(helpers.ERR_UNAUTHORIZED)
		}
		return helpers.EMPTY, errors.New(helpers.ERR_GET_INFO + ":" + err.Error())
	}

	// Validate password
	result := helpers.ComparePass([]byte(users[0].Password), []byte(payload.Password))
	if !result {
		return helpers.EMPTY, errors.New(helpers.ERR_INVALID_PASS)
	}

	// Generate token
	token := helpers.GenerateToken(users[0].Email, users[0].Subscription)

	return token, nil
}

func (s *Service) GetProfiles(email string) (users []model.User, err error) {
	// Get others profile
	users, err = s.UserRepo.SelectUsers(email)
	if err != nil {
		return users, errors.New(helpers.ERR_GET_INFO + ":" + err.Error())
	}

	return users, nil
}

func (s *Service) Swipe(email, idTarget, action, subscription string) error {
	// Get info detail of id Target
	key := fmt.Sprintf("%s:%s", email, idTarget)

	// Check exist target's profile
	exist, err := s.UserRepo.Exist(key)
	if err != nil {
		return errors.New(helpers.ERR_CHECK_EXIST + ":" + err.Error())
	}

	if exist == 1 {
		return errors.New(helpers.ERR_SWIPE)
	}

	// Check total swipe
	total, err := s.UserRepo.GetKeys(key, email)
	if err != nil {
		return errors.New(helpers.ERR_CHECK_TOTAL + ":" + err.Error())
	}

	if total >= 10 && subscription == helpers.SUBSCRIPTION_FREE {
		return errors.New(helpers.ERR_LIMIT)
	}

	// Set Key
	err = s.UserRepo.SetKey(key, action)
	if err != nil {
		return errors.New(helpers.ERR_SET_KEY + ":" + err.Error())
	}

	return nil
}

func (s *Service) Purchase(email string) error {
	// Check user exist
	_, err := s.UserRepo.SelectByCondition(&model.User{
		Email: email,
	})
	if err != nil {
		if err == sql.ErrNoRows {
			return errors.New(helpers.ERR_NOT_FOUND)
		}

		return errors.New(helpers.ERR_GET_INFO + ":" + err.Error())
	}

	// Update subscription
	err = s.UserRepo.UpdateProfile(model.User{
		Email:        email,
		Subscription: helpers.SUBSCRIPTION_PREMIUM,
	})
	if err != nil {
		return errors.New(helpers.ERR_SUBSCRIPTION + ":" + err.Error())
	}

	return nil
}
