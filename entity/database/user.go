package database

type User struct {
	ID        uint   `gorm:"primary_key" json:"id"`
	Username  string `gorm:"column:username" json:"username"`
	Email     string `gorm:"column:email" json:"email"`
	Password  string `gorm:"column:password" json:"password"`
	UserType  string `gorm:"column:user_type"`
	Available bool   `gorm:"column:available"`
}

func (User) TableName() string {
	return "users"
}
