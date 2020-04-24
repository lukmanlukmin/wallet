package database

import (
	"github.com/jinzhu/gorm"
	dbEntity "github.com/lukmanlukmin/wallet/entity/database"
	connection "github.com/lukmanlukmin/wallet/util/connection"
)

type UserBalanceHistoryRepository struct {
	DB gorm.DB
}

func UserBalanceHistoryRepositoryHandler() *UserBalanceHistoryRepository {
	return &UserBalanceHistoryRepository{DB: *connection.GetConnection()}
}

type UserBalanceHistoryRepositoryInterface interface {
	InsertBalanceHistoryData(balanceData *dbEntity.UserBalanceHistories) error
}

func (repository *UserBalanceHistoryRepository) InsertBalanceHistoryData(balanceHistoryata *dbEntity.UserBalanceHistories) error {
	query := repository.DB.Table("user_balance")
	query = query.Create(balanceHistoryata)
	return query.Error
}
