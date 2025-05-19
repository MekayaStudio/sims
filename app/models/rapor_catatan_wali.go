package models

import "github.com/goravel/framework/database/orm"

type RaporCatatanWali struct {
	orm.Model
	IdCatatanWali uint   `gorm:"column:id_catatan_wali;primaryKey"`
	IdTp         int    `gorm:"column:id_tp"`
	IdSmt        int    `gorm:"column:id_smt"`
	IdKelas      int    `gorm:"column:id_kelas"`
	IdSiswa      int    `gorm:"column:id_siswa"`
	Nilai        string `gorm:"column:nilai"`
	Deskripsi    string `gorm:"column:deskripsi"`
}
