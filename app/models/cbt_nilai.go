package models

import "github.com/goravel/framework/database/orm"

type CbtNilai struct {
	orm.Model
	PgBenar        int    `gorm:"column:pg_benar"`
	PgNilai        string `gorm:"column:pg_nilai"`
	EssaiNilai     string `gorm:"column:essai_nilai"`
	IdSiswa        int    `gorm:"column:id_siswa"`
	IdJadwal       int    `gorm:"column:id_jadwal"`
	KompleksNilai  string `gorm:"column:kompleks_nilai"`
	JodohkanNilai  string `gorm:"column:jodohkan_nilai"`
	IsianNilai     string `gorm:"column:isian_nilai"`
	Dikoreksi      int    `gorm:"column:dikoreksi"`
	TimeCreate     string `gorm:"column:time_create"`
}
