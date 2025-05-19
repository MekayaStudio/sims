package models

import "github.com/goravel/framework/database/orm"

type RaporDataFisik struct {
	orm.Model
	IdFisik    uint   `gorm:"column:id_fisik;primaryKey"`
	IdTp       int    `gorm:"column:id_tp"`
	IdSmt      int    `gorm:"column:id_smt"`
	IdKelas    int    `gorm:"column:id_kelas"`
	Jenis      int    `gorm:"column:jenis"`
	Kode       int    `gorm:"column:kode"`
	Deskripsi  string `gorm:"column:deskripsi"`
}
