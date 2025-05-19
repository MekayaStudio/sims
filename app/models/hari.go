package models

import "github.com/goravel/framework/database/orm"

type Hari struct {
	orm.Model
	IdHri   uint   `gorm:"column:id_hri;primaryKey"`
	NamaHri string `gorm:"column:nama_hri"`
}
