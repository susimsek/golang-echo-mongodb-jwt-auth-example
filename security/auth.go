package security

import (
	"golang-echo-mongodb-jwt-auth-example/model"
	"golang-echo-mongodb-jwt-auth-example/repository"
	"golang-echo-mongodb-jwt-auth-example/util"
)

type AuthValidator struct {
	userRepository repository.UserRepository
}

func NewAuthValidator(userRepository repository.UserRepository) *AuthValidator {
	return &AuthValidator{userRepository: userRepository}
}

func (authValidator *AuthValidator) ValidateCredentials(username, password string) (*model.User, bool) {
	user, err := authValidator.userRepository.FindByEmail(username)
	if err != nil || util.VerifyPassword(user.Password, password) != nil {
		return nil, false
	}
	return user, true
}
