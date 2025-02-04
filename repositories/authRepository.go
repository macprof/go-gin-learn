package repositories

import (
	"context"
	"go-gin-workshop/entities"

	"gorm.io/gorm"
)

type AuthRepository interface {
	LoginRepository()
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

func (r *authRepository) LoginRepository() {

}
func (r *authRepository) RegisterRepository(ctx context.Context, modelEntity entities.UserEntity) (entities.UserEntity, error) {
	if err := r.db.Create(&modelEntity).Error; err != nil {
		return entities.UserEntity{}, err
	}
	return modelEntity, nil
}
