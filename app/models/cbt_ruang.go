package models

import "github.com/goravel/framework/database/orm"

type CbtRuang struct {
	orm.Model
	IdRuang    uint   `gorm:"column:id_ruang;primaryKey"`
	NamaRuang  string `gorm:"column:nama_ruang"`
	KodeRuang  string `gorm:"column:kode_ruang"`
}
