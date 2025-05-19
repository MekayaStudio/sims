package models

import "github.com/goravel/framework/database/orm"

type RaporNilaiAkhir struct {
	orm.Model
	IdNilaiAkhir uint   `gorm:"column:id_nilai_akhir;primaryKey"`
	IdMapel      int    `gorm:"column:id_mapel"`
	IdKelas      int    `gorm:"column:id_kelas"`
	IdSiswa      int    `gorm:"column:id_siswa"`
	IdTp         int    `gorm:"column:id_tp"`
	IdSmt        int    `gorm:"column:id_smt"`
	Nilai        int    `gorm:"column:nilai"`
	Akhir        int    `gorm:"column:akhir"`
	Predikat     string `gorm:"column:predikat"`
}
