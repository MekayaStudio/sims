package models

import "github.com/goravel/framework/database/orm"

type ApiToken struct {
	orm.Model
	IdApi     uint   `gorm:"column:id_api;primaryKey"`
	Timestamp string `gorm:"column:timestamp"`
	IdUser    int    `gorm:"column:id_user"`
	Address   string `gorm:"column:address"`
	Agent     string `gorm:"column:agent"`
	Device    string `gorm:"column:device"`
	Token     string `gorm:"column:token"`
}
