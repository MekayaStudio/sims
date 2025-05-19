package models

import "github.com/goravel/framework/database/orm"

type Bulan struct {
	orm.Model
	IdBln   uint   `gorm:"column:id_bln;primaryKey"`
	NamaBln string `gorm:"column:nama_bln"`
}
