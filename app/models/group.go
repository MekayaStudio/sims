package models

import "github.com/goravel/framework/database/orm"

type Group struct {
	orm.Model
	ID          uint   `gorm:"column:id;primaryKey"`
	Name        string `gorm:"column:name"`
	Description string `gorm:"column:description"`
}
