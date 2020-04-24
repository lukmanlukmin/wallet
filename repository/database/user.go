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
	GetUserByEmail(email string, userData *dbEntity.User) error
	GetUserByEmailPassword(email string, password string, userData *dbEntity.User) error
	GetUserList(limit int, offset int) ([]dbEntity.User, error)
	InsertUser(userData *dbEntity.User) error
	UpdateUser(id int, userData *dbEntity.User) error
	DeleteUser(id int) error
}

func (repository *UserRepository) GetUserByEmail(email string, userData *dbEntity.User) error {
	query := repository.DB.Table("users")
	query = query.Where("email=?", email)
	query = query.First(userData)
	return query.Error
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
func (repository *UserRepository) GetUserList(limit int, offset int) ([]dbEntity.User, error) {
	users := []dbEntity.User{}
	query := repository.DB.Table("users")
	query = query.Limit(limit).Offset(offset)
	query = query.Find(&users)
	return users, query.Error
}

func (repository *UserRepository) InsertUser(userData *dbEntity.User) error {
	query := repository.DB.Table("users")
	query = query.Create(userData)
	return query.Error
}

func (repository *UserRepository) UpdateUser(id int, userData *dbEntity.User) error {
	query := repository.DB.Table("users")
	query = query.Where("id=?", id)
	query = query.Updates(userData)
	return query.Error
}

func (repository *UserRepository) DeleteUser(id int) error {
	query := repository.DB.Table("users")
	query = query.Where("id=?", id)
	query = query.Delete(&dbEntity.User{})
	return query.Error
}
