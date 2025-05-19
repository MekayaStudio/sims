package models

import "github.com/goravel/framework/database/orm"

type RaporDataSikap struct {
	orm.Model
	IdSikap   uint   `gorm:"column:id_sikap;primaryKey"`
	IdTp      int    `gorm:"column:id_tp"`
	IdSmt     int    `gorm:"column:id_smt"`
	IdKelas   int    `gorm:"column:id_kelas"`
	Jenis     int    `gorm:"column:jenis"`
	Kode      int    `gorm:"column:kode"`
	Sikap     string `gorm:"column:sikap"`
}
