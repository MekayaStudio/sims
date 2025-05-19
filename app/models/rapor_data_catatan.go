package models

import "github.com/goravel/framework/database/orm"

type RaporDataCatatan struct {
	orm.Model
	IdCatatan   uint   `gorm:"column:id_catatan;primaryKey"`
	IdTp        int    `gorm:"column:id_tp"`
	IdSmt       int    `gorm:"column:id_smt"`
	IdKelas     int    `gorm:"column:id_kelas"`
	Jenis       int    `gorm:"column:jenis"`
	Kode        int    `gorm:"column:kode"`
	Deskripsi   string `gorm:"column:deskripsi"`
	Rank        string `gorm:"column:rank"`
}
