package database

import (
	"sync"

	"github.com/jinzhu/gorm"
	dbEntity "github.com/lukmanlukmin/wallet/entity/database"
	connection "github.com/lukmanlukmin/wallet/util/helper/mysqlconnection"
)

type UserRepository struct {
	DB gorm.DB
}

func UserRepositoryHandler() *UserRepository {
	return &UserRepository{DB: *connection.GetConnection()}
}

type UserRepositoryInterface interface {
	GetUserByID(id int, userData *dbEntity.User, wg *sync.WaitGroup) error
}

func (repository *UserRepository) GetUserByID(id int, userData *dbEntity.User, wg *sync.WaitGroup) error {
	query := repository.DB.Preload("UserStatus")
	query = query.Where("id=?", id)
	query = query.First(userData)
	wg.Done()
	return query.Error
}

func (repository *UserRepository) GetUserByEmailPassword(email string, password string, userData *dbEntity.User, wg *sync.WaitGroup) error {
	query := repository.DB.Preload("UserStatus")
	query = query.Where("email=? AND password=?", email, password)
	query = query.First(userData)
	wg.Done()
	return query.Error
}
