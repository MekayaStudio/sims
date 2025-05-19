package models

import "github.com/goravel/framework/database/orm"

type CbtToken struct {
	orm.Model
	IdToken  uint   `gorm:"column:id_token;primaryKey"`
	Token    string `gorm:"column:token"`
	Auto     int    `gorm:"column:auto"`
	Jarak    int    `gorm:"column:jarak"`
	Updated  string `gorm:"column:updated"`
}
