package models

import "github.com/goravel/framework/database/orm"

type MasterKelas struct {
	orm.Model
	IdKelas      uint   `gorm:"column:id_kelas;primaryKey"`
	IdTp        int    `gorm:"column:id_tp"`
	IdSmt       int    `gorm:"column:id_smt"`
	NamaKelas   string `gorm:"column:nama_kelas"`
	KodeKelas   string `gorm:"column:kode_kelas"`
	JurusanId   int    `gorm:"column:jurusan_id"`
	LevelId     int    `gorm:"column:level_id"`
	GuruId      int    `gorm:"column:guru_id"`
	SiswaId     int    `gorm:"column:siswa_id"`
	JumlahSiswa string `gorm:"column:jumlah_siswa"`
	SetSiswa    string `gorm:"column:set_siswa"`
}
