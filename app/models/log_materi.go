package models

import "github.com/goravel/framework/database/orm"

type LogMateri struct {
	orm.Model
	IdLog      string `gorm:"column:id_log;primaryKey"`
	LogTime    string `gorm:"column:log_time"`
	IdSiswa    int    `gorm:"column:id_siswa"`
	JamKe      int    `gorm:"column:jam_ke"`
	IdMateri   string `gorm:"column:id_materi"`
	IdMapel    int    `gorm:"column:id_mapel"`
	LogType    int    `gorm:"column:log_type"`
	LogDesc    string `gorm:"column:log_desc"`
	Text       string `gorm:"column:text"`
	File       string `gorm:"column:file"`
	Nilai      string `gorm:"column:nilai"`
	Catatan    string `gorm:"column:catatan"`
	Address    string `gorm:"column:address"`
	Agent      string `gorm:"column:agent"`
	Device     string `gorm:"column:device"`
	FinishTime string `gorm:"column:finish_time"`
}
