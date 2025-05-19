package models

import "github.com/goravel/framework/database/orm"

type CbtSesi struct {
	orm.Model
	IdSesi     uint   `gorm:"column:id_sesi;primaryKey"`
	NamaSesi   string `gorm:"column:nama_sesi"`
	KodeSesi   string `gorm:"column:kode_sesi"`
	WaktuMulai string `gorm:"column:waktu_mulai"`
	WaktuAkhir string `gorm:"column:waktu_akhir"`
	Aktif      int    `gorm:"column:aktif"`
}
