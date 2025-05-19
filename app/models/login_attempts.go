package models

type LoginAttempts struct {
	ID         int    `gorm:"column:id;primaryKey"`
	IPAddress  string `gorm:"column:ip_address"`
	Login      string `gorm:"column:login"`
	Time       int    `gorm:"column:time"`
}
