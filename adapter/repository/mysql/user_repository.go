package mysql

import (
	"context"

	"github.com/tocura/go-jwt-authentication/adapter/repository/mysql/entity"
	"github.com/tocura/go-jwt-authentication/core/model"
	"github.com/tocura/go-jwt-authentication/core/port"
	"github.com/tocura/go-jwt-authentication/pkg/log"
	"gorm.io/gorm"
)

const usersTable = "users"

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) port.UserRepository {
	return &userRepository{
		db: db,
	}
}

func (u *userRepository) Create(ctx context.Context, user model.User) (*model.User, error) {
	userEntity := entity.MapToUserEntity(user)
	userEntity.SetCreatedAt()
	userEntity.SetID()

	result := u.db.WithContext(ctx).
		Table(usersTable).
		Create(&userEntity)

	if result.Error != nil {
		log.Error(ctx, "error to create user in database", result.Error)
		return nil, result.Error
	}

	log.Info(ctx, "user saved in database with success")
	return userEntity.MapToUserModel(), nil
}

func (u *userRepository) Login(ctx context.Context, login model.Login) (*model.User, error) {
	loginEntity := entity.MapToLoginEntity(login)
	var userEntity entity.User

	result := u.db.WithContext(ctx).
		Table(usersTable).
		Where("email = ? AND password = ?", loginEntity.Email, login.Password).
		First(&userEntity)

	if result.Error != nil {
		log.Error(ctx, "user not find", result.Error)
		return nil, result.Error
	}

	log.Info(ctx, "user credentials retrieved from database with success")
	return userEntity.MapToUserModel(), nil
}
