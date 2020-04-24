package database

import (
	"github.com/jinzhu/gorm"
	dbEntity "github.com/lukmanlukmin/wallet/entity/database"
	connection "github.com/lukmanlukmin/wallet/util/connection"
)

type UserBalanceRepository struct {
	DB gorm.DB
}

func UserBalanceRepositoryHandler() *UserBalanceRepository {
	return &UserBalanceRepository{DB: *connection.GetConnection()}
}

type UserBalanceRepositoryInterface interface {
	GetCurrentBalance(id int, balanceData *dbEntity.UserBalances) error
	InsertBalanceData(balanceData *dbEntity.UserBalances) error
}

func (repository *UserBalanceRepository) GetCurrentBalance(id int, balanceData *dbEntity.UserBalances) error {
	query := repository.DB.Table("user_balances")
	query = query.Where("user_id=?", id)
	query = query.Order("id desc")
	query = query.First(&balanceData)
	return query.Error
}

func (repository *UserBalanceRepository) InsertBalanceData(balanceData *dbEntity.UserBalances) error {
	query := repository.DB.Table("user_balances")
	query = query.Create(balanceData)
	return query.Error
}
