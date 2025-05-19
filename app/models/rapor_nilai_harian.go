package models

import "github.com/goravel/framework/database/orm"

type RaporNilaiHarian struct {
	orm.Model
	IdNilaiHarian uint   `gorm:"column:id_nilai_harian;primaryKey"`
	IdSiswa       int    `gorm:"column:id_siswa"`
	IdMapel       int    `gorm:"column:id_mapel"`
	IdKelas       int    `gorm:"column:id_kelas"`
	IdTp          int    `gorm:"column:id_tp"`
	IdSmt         int    `gorm:"column:id_smt"`
	P1            string `gorm:"column:p1"`
	P2            string `gorm:"column:p2"`
	P3            string `gorm:"column:p3"`
	P4            string `gorm:"column:p4"`
	P5            string `gorm:"column:p5"`
	P6            string `gorm:"column:p6"`
	P7            string `gorm:"column:p7"`
	P8            string `gorm:"column:p8"`
	PRataRata     string `gorm:"column:p_rata_rata"`
	PPredikat     string `gorm:"column:p_predikat"`
	PDeskripsi    string `gorm:"column:p_deskripsi"`
	K1            string `gorm:"column:k1"`
	K2            string `gorm:"column:k2"`
	K3            string `gorm:"column:k3"`
	K4            string `gorm:"column:k4"`
	K5            string `gorm:"column:k5"`
}
