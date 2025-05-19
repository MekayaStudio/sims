package models

import "github.com/goravel/framework/database/orm"

type CbtDurasiSiswa struct {
	orm.Model
	IdDurasi   uint   `gorm:"column:id_durasi;primaryKey"`
	IdSiswa    int    `gorm:"column:id_siswa"`
	IdJadwal   int    `gorm:"column:id_jadwal"`
	Status     int    `gorm:"column:status"`
	LamaUjian  string `gorm:"column:lama_ujian"`
	Mulai      string `gorm:"column:mulai"`
	Selesai    string `gorm:"column:selesai"`
	Reset      int    `gorm:"column:reset"`
	TimeCreate string `gorm:"column:time_create"`
}
