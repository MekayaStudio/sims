package models

import "github.com/goravel/framework/database/orm"

// CbtJadwal represents the cbt_jadwal table
// Only main fields are included, add more as needed
// Adjust types if needed for your ORM

type CbtJadwal struct {
	orm.Model
	IdJadwal      uint   `gorm:"column:id_jadwal;primaryKey"`
	IdTp          string `gorm:"column:id_tp"`
	IdSmt         string `gorm:"column:id_smt"`
	IdBank        int    `gorm:"column:id_bank"`
	IdJenis       int    `gorm:"column:id_jenis"`
	TglMulai      string `gorm:"column:tgl_mulai"`
	TglSelesai    string `gorm:"column:tgl_selesai"`
	DurasiUjian   int    `gorm:"column:durasi_ujian"`
	Pengawas      string `gorm:"column:pengawas"`
	AcakSoal      int    `gorm:"column:acak_soal"`
	AcakOpsi      int    `gorm:"column:acak_opsi"`
	HasilTampil   int    `gorm:"column:hasil_tampil"`
	Token         int    `gorm:"column:token"`
	Status        int    `gorm:"column:status"`
	Ulang         int    `gorm:"column:ulang"`
	ResetLogin    int    `gorm:"column:reset_login"`
	Rekap         int    `gorm:"column:rekap"`
	JamKe         int    `gorm:"column:jam_ke"`
	Jarak         int    `gorm:"column:jarak"`
	TimeCreate    string `gorm:"column:time_create"`
}
