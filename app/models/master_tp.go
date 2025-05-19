package models

import "github.com/goravel/framework/database/orm"

type MasterTp struct {
	orm.Model
	IdTp   uint   `gorm:"column:id_tp;primaryKey"`
	Tahun  string `gorm:"column:tahun"`
	Active int    `gorm:"column:active"`
}
