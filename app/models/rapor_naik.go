package models

import "github.com/goravel/framework/database/orm"

type RaporNaik struct {
	orm.Model
	IdNaik   int `gorm:"column:id_naik;primaryKey"`
	IdTp     int `gorm:"column:id_tp"`
	IdSmt    int `gorm:"column:id_smt"`
	IdSiswa  int `gorm:"column:id_siswa"`
	Naik     int `gorm:"column:naik"`
}
