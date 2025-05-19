package models

import "github.com/goravel/framework/database/orm"

type CbtNomorPeserta struct {
	orm.Model
	IdNomor      uint   `gorm:"column:id_nomor;primaryKey"`
	IdSiswa      int    `gorm:"column:id_siswa"`
	IdTp         int    `gorm:"column:id_tp"`
	IdSmt        int    `gorm:"column:id_smt"`
	NomorPeserta string `gorm:"column:nomor_peserta"`
}
