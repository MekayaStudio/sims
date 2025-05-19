package models

import "github.com/goravel/framework/database/orm"

type RaporNilaiEkstra struct {
	orm.Model
	IdNilaiEkstra uint   `gorm:"column:id_nilai_ekstra;primaryKey"`
	IdEkstra      int    `gorm:"column:id_ekstra"`
	IdKelas       int    `gorm:"column:id_kelas"`
	IdSiswa       int    `gorm:"column:id_siswa"`
	IdTp          int    `gorm:"column:id_tp"`
	IdSmt         int    `gorm:"column:id_smt"`
	Nilai         int    `gorm:"column:nilai"`
	Predikat      string `gorm:"column:predikat"`
	Deskripsi     string `gorm:"column:deskripsi"`
}
