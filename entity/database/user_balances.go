package database

type UserBalances struct {
	ID              uint `gorm:"primary_key" json:"id"`
	UserID          uint `gorm:"column:user_id" json:"user_id"`
	Balance         uint `gorm:"column:balance" json:"balance"`
	BalanceAchieves uint `gorm:"column:balance_achieves" json:"balance_achieves"`
}

func (UserBalances) TableName() string {
	return "user_balances"
}
