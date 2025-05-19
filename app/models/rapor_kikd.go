package models

import "github.com/goravel/framework/database/orm"

type RaporKikd struct {
	orm.Model
	IdKikd        uint   `gorm:"column:id_kikd;primaryKey"`
	IdMapelKelas  int    `gorm:"column:id_mapel_kelas"`
	Aspek         int    `gorm:"column:aspek"`
	IdTp          int    `gorm:"column:id_tp"`
	IdSmt         int    `gorm:"column:id_smt"`
	MateriKikd    string `gorm:"column:materi_kikd"`
}
