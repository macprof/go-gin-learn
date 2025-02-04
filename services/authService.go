package services

import (
	"context"
	"go-gin-workshop/dto"
	"go-gin-workshop/entities"
	"go-gin-workshop/repositories"
)

type AuthService interface {
	LoginService()
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

func (c *authService) LoginService() {

}

func (c *authService) RegisterService(ctx context.Context, req dto.UserDto) (entities.UserEntity, error) {

	modelUser := entities.UserEntity{
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Email:     req.Email,
		Phone:     req.Phone,
		Password:  req.Password,
		Platform:  req.Platform,
	}

	result, err := c.authRepository.RegisterRepository(ctx, modelUser)

	return result, err
}
