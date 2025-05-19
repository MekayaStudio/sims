package models

import "github.com/goravel/framework/database/orm"

type LogUjian struct {
	orm.Model
	IdLog      uint   `gorm:"column:id_log;primaryKey"`
	LogTime    string `gorm:"column:log_time"`
	IdSiswa    int    `gorm:"column:id_siswa"`
	IdJadwal   int    `gorm:"column:id_jadwal"`
	LogType    int    `gorm:"column:log_type"`
	LogDesc    string `gorm:"column:log_desc"`
	Address    string `gorm:"column:address"`
	Agent      string `gorm:"column:agent"`
	Device     string `gorm:"column:device"`
	Reset      int    `gorm:"column:reset"`
	FinishTime string `gorm:"column:finish_time"`
}
