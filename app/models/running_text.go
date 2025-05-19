package models

import "github.com/goravel/framework/database/orm"

type RunningText struct {
	orm.Model
	IdText uint   `gorm:"column:id_text;primaryKey"`
	Text   string `gorm:"column:text"`
}
