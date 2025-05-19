package models

import "github.com/goravel/framework/database/orm"

type MasterJurusan struct {
	orm.Model
	IdJurusan       uint   `gorm:"column:id_jurusan;primaryKey"`
	NamaJurusan     string `gorm:"column:nama_jurusan"`
	KodeJurusan     string `gorm:"column:kode_jurusan"`
	MapelPeminatan  string `gorm:"column:mapel_peminatan"`
	Status          int    `gorm:"column:status"`
	Deletable       int    `gorm:"column:deletable"`
}
