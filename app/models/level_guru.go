package models

import "github.com/goravel/framework/database/orm"

type LevelGuru struct {
	orm.Model
	IdLevel uint   `gorm:"column:id_level;primaryKey"`
	Level   string `gorm:"column:level"`
}
