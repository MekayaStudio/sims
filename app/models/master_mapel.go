package models

import "github.com/goravel/framework/database/orm"

type MasterMapel struct {
	orm.Model
	IdMapel      uint   `gorm:"column:id_mapel;primaryKey"`
	NamaMapel    string `gorm:"column:nama_mapel"`
	Kode         string `gorm:"column:kode"`
	Kelompok     string `gorm:"column:kelompok"`
	BobotP       int    `gorm:"column:bobot_p"`
	BobotK       int    `gorm:"column:bobot_k"`
	Jenjang      int    `gorm:"column:jenjang"`
	Urutan       int    `gorm:"column:urutan"`
	Status       int    `gorm:"column:status"`
	Deletable    int    `gorm:"column:deletable"`
	UrutanTampil int    `gorm:"column:urutan_tampil"`
}
