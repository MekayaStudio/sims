package models

import "github.com/goravel/framework/database/orm"

type CbtPengawas struct {
	orm.Model
	IdPengawas string `gorm:"column:id_pengawas;primaryKey"`
	IdJadwal   string `gorm:"column:id_jadwal"`
	IdTp       int    `gorm:"column:id_tp"`
	IdSmt      int    `gorm:"column:id_smt"`
	IdRuang    string `gorm:"column:id_ruang"`
	IdSesi     string `gorm:"column:id_sesi"`
	IdGuru     string `gorm:"column:id_guru"`
}
