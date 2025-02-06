package repositories

import (
	"context"
	"errors"
	"go-gin-workshop/dto"
	"go-gin-workshop/entities"

	"gorm.io/gorm"
)

type AuthRepository interface {
	LoginRepository(ctx context.Context, req dto.LoginDto) (entities.UserEntity, error)

	RegisterRepository(ctx context.Context, modelEntity entities.UserEntity) (entities.UserEntity, error)
}

type authRepository struct {
	db *gorm.DB
}

func NewAuthRepository(db *gorm.DB) AuthRepository {
	return &authRepository{
		db: db,
	}
}

func (r *authRepository) LoginRepository(ctx context.Context, req dto.LoginDto) (entities.UserEntity, error) {
	var existingUser entities.UserEntity

	result := r.db.Where("email = ?", req.Email).First(&existingUser)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return entities.UserEntity{}, errors.New("email does not exist")
		}
		return entities.UserEntity{}, result.Error
	}

	return existingUser, nil

}
func (r *authRepository) RegisterRepository(ctx context.Context, modelEntity entities.UserEntity) (entities.UserEntity, error) {
	var existingUser entities.UserEntity

	err := r.db.Where("email = ?", modelEntity.Email).First(&existingUser).Error
	if err == nil {
		return entities.UserEntity{}, errors.New("Email already exists")
	}
	if err := r.db.Create(&modelEntity).Error; err != nil {
		return entities.UserEntity{}, err
	}
	return modelEntity, nil
}
