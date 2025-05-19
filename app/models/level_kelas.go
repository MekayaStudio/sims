package models

import "github.com/goravel/framework/database/orm"

type LevelKelas struct {
	orm.Model
	IdLevel int    `gorm:"column:id_level"`
	Level   string `gorm:"column:level"`
}
