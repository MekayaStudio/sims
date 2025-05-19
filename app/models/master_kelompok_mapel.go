package models

import "github.com/goravel/framework/database/orm"

type MasterKelompokMapel struct {
	orm.Model
	IdKelMapel     uint   `gorm:"column:id_kel_mapel;primaryKey"`
	KodeKelMapel   string `gorm:"column:kode_kel_mapel"`
	NamaKelMapel   string `gorm:"column:nama_kel_mapel"`
	Kategori       string `gorm:"column:kategori"`
	IdParent       int    `gorm:"column:id_parent"`
}
