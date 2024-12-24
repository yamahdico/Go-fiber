package models

type User struct {
	ID       uint   `gorm:"primaryKey;autoIncrement"`
	Username string `json:"username" gorm:"unique"`
	Password string `json:"password" gorm:"type:varchar(10)"`
	Email    string `json:"email" gorm:"unique"`
}
