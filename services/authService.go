package services

import (
	"context"
	"errors"
	"go-gin-workshop/dto"
	"go-gin-workshop/entities"
	"go-gin-workshop/repositories"
	"go-gin-workshop/utils"
	"regexp"
)

type AuthService interface {
	LoginService(ctx context.Context, req dto.LoginDto) (entities.UserEntity, error)
	RegisterService(ctx context.Context, req dto.UserDto) (entities.UserEntity, error)
}

type authService struct {
	authRepository repositories.AuthRepository
}

func NewAuthService(authRepository repositories.AuthRepository) AuthService {
	return &authService{
		authRepository: authRepository,
	}
}

func (c *authService) LoginService(ctx context.Context, req dto.LoginDto) (entities.UserEntity, error) {

	if !isValidEmail(req.Email) {
		return entities.UserEntity{}, errors.New("Invalid email format")
	}
	existingUser, err := c.authRepository.LoginRepository(ctx, req)
	if err != nil {
		return entities.UserEntity{}, err
	}
	passwordDehash := utils.CheckPassword(existingUser.Password, req.Password)
	if passwordDehash == false {
		return entities.UserEntity{}, errors.New("Invalid password")
	}
	return existingUser, err
}

var emailRegex = regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)

func (c *authService) RegisterService(ctx context.Context, req dto.UserDto) (entities.UserEntity, error) {

	if !isValidEmail(req.Email) {
		return entities.UserEntity{}, errors.New("Invalid email format")
	}
	passwordHash, err := utils.HashPassword(req.Password)
	modelUser := entities.UserEntity{
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Email:     req.Email,
		Phone:     req.Phone,
		Password:  passwordHash,
		Platform:  req.Platform,
	}
	result, err := c.authRepository.RegisterRepository(ctx, modelUser)
	if err != nil {
		return entities.UserEntity{}, err
	}

	return result, nil

}

func isValidEmail(email string) bool {
	return emailRegex.MatchString(email)
}
