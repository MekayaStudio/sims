package models

import "github.com/goravel/framework/database/orm"

type RaporNilaiPts struct {
	orm.Model
	IdNilaiPts uint   `gorm:"column:id_nilai_pts;primaryKey"`
	IdMapel    int    `gorm:"column:id_mapel"`
	IdKelas    int    `gorm:"column:id_kelas"`
	IdSiswa    int    `gorm:"column:id_siswa"`
	IdTp       int    `gorm:"column:id_tp"`
	IdSmt      int    `gorm:"column:id_smt"`
	Nilai      int    `gorm:"column:nilai"`
	Predikat   string `gorm:"column:predikat"`
}
