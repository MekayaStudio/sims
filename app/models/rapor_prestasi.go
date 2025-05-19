package models

import "github.com/goravel/framework/database/orm"

type RaporPrestasi struct {
	orm.Model
	IdRanking   uint   `gorm:"column:id_ranking;primaryKey"`
	IdKelas     int    `gorm:"column:id_kelas"`
	IdSiswa     int    `gorm:"column:id_siswa"`
	IdTp        int    `gorm:"column:id_tp"`
	IdSmt       int    `gorm:"column:id_smt"`
	Ranking     int    `gorm:"column:ranking"`
	Deskripsi   string `gorm:"column:deskripsi"`
	P1          string `gorm:"column:p1"`
	P1Desk      string `gorm:"column:p1_desk"`
	P2          string `gorm:"column:p2"`
	P2Desk      string `gorm:"column:p2_desk"`
	P3          string `gorm:"column:p3"`
	P3Desk      string `gorm:"column:p3_desk"`
}
