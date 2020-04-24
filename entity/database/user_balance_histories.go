package database

import "time"

type UserBalanceHistories struct {
	ID            uint       `gorm:"primary_key" json:"id"`
	UserBalanceID uint       `gorm:"column:user_balance_id" json:"user_balance_id"`
	BalanceBefore uint       `gorm:"column:balance_before" json:"balance_before"`
	BalanceAfter  uint       `gorm:"column:balance_after" json:"balance_after"`
	Activity      string     `gorm:"column:activity" json:"activity"`
	TypeActivity  string     `gorm:"column:type_activity" json:"type_activity"`
	IP            string     `gorm:"column:ip" json:"ip"`
	UserAgent     string     `gorm:"column:user_agent" json:"user_agent"`
	CreatedAt     *time.Time `gorm:"column:created_at" json:"created_at"`
}

func (UserBalanceHistories) TableName() string {
	return "user_balance_histories"
}
