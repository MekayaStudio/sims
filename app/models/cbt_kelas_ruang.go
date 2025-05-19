package models

import "github.com/goravel/framework/database/orm"

type CbtKelasRuang struct {
	orm.Model
	IdKelasRuang string `gorm:"column:id_kelas_ruang;primaryKey"`
	IdKelas      int    `gorm:"column:id_kelas"`
	IdRuang      int    `gorm:"column:id_ruang"`
	IdSesi       int    `gorm:"column:id_sesi"`
	IdTp         int    `gorm:"column:id_tp"`
	IdSmt        int    `gorm:"column:id_smt"`
	SetSiswa     int    `gorm:"column:set_siswa"`
}
