package model

type Admin struct {
	ID       int    `gorm:"primaryKey"`
	Email    string `gorm:"uniqueIndex"`
	Password string `gorm:"size:255"`
}
