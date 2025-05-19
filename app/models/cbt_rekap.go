package models

import "github.com/goravel/framework/database/orm"

type CbtRekap struct {
	orm.Model
	IdRekap      uint   `gorm:"column:id_rekap;primaryKey"`
	IdTp         int    `gorm:"column:id_tp"`
	Tp           string `gorm:"column:tp"`
	IdSmt        int    `gorm:"column:id_smt"`
	Smt          string `gorm:"column:smt"`
	IdJadwal     string `gorm:"column:id_jadwal"`
	IdJenis      string `gorm:"column:id_jenis"`
	KodeJenis    string `gorm:"column:kode_jenis"`
	IdBank       string `gorm:"column:id_bank"`
	BankKelas    string `gorm:"column:bank_kelas"`
	BankKode     string `gorm:"column:bank_kode"`
	BankLevel    int    `gorm:"column:bank_level"`
	IdMapel      string `gorm:"column:id_mapel"`
	NamaMapel    string `gorm:"column:nama_mapel"`
	Kode         string `gorm:"column:kode"`
	TglMulai     string `gorm:"column:tgl_mulai"`
	TglSelesai   string `gorm:"column:tgl_selesai"`
	TampilPg     int    `gorm:"column:tampil_pg"`
	JawabanPg    string `gorm:"column:jawaban_pg"`
	TampilEsai   int    `gorm:"column:tampil_esai"`
	JawabanEsai  string `gorm:"column:jawaban_esai"`
	BobotPg      int    `gorm:"column:bobot_pg"`
}
