package models

import "github.com/goravel/framework/database/orm"

type RaporFisik struct {
	orm.Model
	IdFisik   uint   `gorm:"column:id_fisik;primaryKey"`
	IdKelas   int    `gorm:"column:id_kelas"`
	IdSiswa   int    `gorm:"column:id_siswa"`
	IdTp      int    `gorm:"column:id_tp"`
	IdSmt     int    `gorm:"column:id_smt"`
	Kondisi   string `gorm:"column:kondisi"`
	Tinggi    int    `gorm:"column:tinggi"`
	Berat     int    `gorm:"column:berat"`
}
