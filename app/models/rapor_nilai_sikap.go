package models

import "github.com/goravel/framework/database/orm"

type RaporNilaiSikap struct {
	orm.Model
	IdNilaiSikap uint   `gorm:"column:id_nilai_sikap;primaryKey"`
	IdSiswa      int    `gorm:"column:id_siswa"`
	IdKelas      int    `gorm:"column:id_kelas"`
	IdTp         int    `gorm:"column:id_tp"`
	IdSmt        int    `gorm:"column:id_smt"`
	Jenis        int    `gorm:"column:jenis"`
	Nilai        string `gorm:"column:nilai"`
	Deskripsi    string `gorm:"column:deskripsi"`
}
