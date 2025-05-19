package models

import "github.com/goravel/framework/database/orm"

type CbtSesiSiswa struct {
	orm.Model
	SiswaId int `gorm:"column:siswa_id"`
	KelasId int `gorm:"column:kelas_id"`
	RuangId int `gorm:"column:ruang_id"`
	SesiId  int `gorm:"column:sesi_id"`
	TpId    int `gorm:"column:tp_id"`
	SmtId   int `gorm:"column:smt_id"`
}
