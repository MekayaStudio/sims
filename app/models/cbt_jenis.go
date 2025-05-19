package models

import "github.com/goravel/framework/database/orm"

type CbtJenis struct {
	orm.Model
	NamaJenis string `gorm:"column:nama_jenis"`
	KodeJenis string `gorm:"column:kode_jenis"`
}
