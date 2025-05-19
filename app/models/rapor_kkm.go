package models

import "github.com/goravel/framework/database/orm"

type RaporKkm struct {
	orm.Model
	IdKkm      uint   `gorm:"column:id_kkm;primaryKey"`
	Kkm        int    `gorm:"column:kkm"`
	BobotPh    int    `gorm:"column:bobot_ph"`
	BobotPts   int    `gorm:"column:bobot_pts"`
	BobotPas   int    `gorm:"column:bobot_pas"`
	BobotAbsen int    `gorm:"column:bobot_absen"`
	BebanJam   int    `gorm:"column:beban_jam"`
	IdTp       int    `gorm:"column:id_tp"`
	IdSmt      int    `gorm:"column:id_smt"`
	Jenis      int    `gorm:"column:jenis"`
	IdKelas    int    `gorm:"column:id_kelas"`
	IdMapel    int    `gorm:"column:id_mapel"`
}
