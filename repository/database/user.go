package database

import (
	"github.com/jinzhu/gorm"
	dbEntity "github.com/lukmanlukmin/wallet/entity/database"
	connection "github.com/lukmanlukmin/wallet/util/connection"
)

type UserRepository struct {
	DB gorm.DB
}

func UserRepositoryHandler() *UserRepository {
	return &UserRepository{DB: *connection.GetConnection()}
}

type UserRepositoryInterface interface {
	GetUserByID(id int, userData *dbEntity.User) error
	GetUserByEmailPassword(email string, password string, userData *dbEntity.User) error
}

func (repository *UserRepository) GetUserByID(id int, userData *dbEntity.User) error {
	query := repository.DB.Table("users")
	query = query.Where("id=?", id)
	query = query.First(userData)
	return query.Error
}

func (repository *UserRepository) GetUserByEmailPassword(email string, password string, userData *dbEntity.User) error {
	query := repository.DB.Preload("UserStatus")
	query = query.Where("email=?", email)
	query = query.Where("password=?", password)
	query = query.First(userData)
	return query.Error
}
