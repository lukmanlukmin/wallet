package response

import "time"

type Balances struct {
	ID              uint `json:"id"`
	UserID          uint `json:"user_id"`
	Balance         uint `json:"balance"`
	BalanceAchieves uint `json:"balance_achieves"`
}

type BalanceHistories struct {
	ID            uint       `json:"id"`
	UserBalanceID uint       `json:"user_balance_id"`
	BalanceBefore uint       `json:"balance_before"`
	BalanceAfter  uint       `json:"balance_after"`
	Activity      string     `json:"activity"`
	TypeActivity  string     `json:"type_activity"`
	IP            string     `json:"ip"`
	UserAgent     string     `json:"user_agent"`
	CreatedAt     *time.Time `json:"created_at"`
}
