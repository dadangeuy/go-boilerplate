package model

type User struct {
	ID       uint	`gorm:"primaryKey;autoIncrement;not null"`
	Email    string	`gorm:"unique;size:255;not null"`
	Password string	`gorm:"size:255;not null"`
}

func (User) TableName() string {
	return "users"
}
