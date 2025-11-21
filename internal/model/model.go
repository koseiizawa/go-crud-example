package model

type UserModel struct {
	ID       int    `gorm:"primaryKey"`
	Name     string `gorm:"size:255"`
	Email    string `gorm:"uniqueIndex"`
	Password string `gorm:"size:255"`
}
