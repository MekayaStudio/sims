package models

import "github.com/goravel/framework/database/orm"

type LoginAttempt struct {
	orm.Model
	ID        uint   `gorm:"column:id;primaryKey"`
	IpAddress string `gorm:"column:ip_address"`
	Login     string `gorm:"column:login"`
	Time      uint   `gorm:"column:time"`
}
