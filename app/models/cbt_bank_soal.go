package models

import "github.com/goravel/framework/database/orm"

type CbtBankSoal struct {
	orm.Model
	IdBank         uint   `gorm:"column:id_bank;primaryKey"`
	BankJenisId    int    `gorm:"column:bank_jenis_id"`
	BankKode       string `gorm:"column:bank_kode"`
	BankLevel      string `gorm:"column:bank_level"`
	BankKelas      string `gorm:"column:bank_kelas"`
	BankMapelId    int    `gorm:"column:bank_mapel_id"`
	BankJurusanId  int    `gorm:"column:bank_jurusan_id"`
	BankGuruId     int    `gorm:"column:bank_guru_id"`
	BankNama       string `gorm:"column:bank_nama"`
	Kkm            int    `gorm:"column:kkm"`
	JmlSoal        int    `gorm:"column:jml_soal"`
	JmlEsai        int    `gorm:"column:jml_esai"`
	TampilPg       int    `gorm:"column:tampil_pg"`
	TampilEsai     int    `gorm:"column:tampil_esai"`
	BobotPg        int    `gorm:"column:bobot_pg"`
	BobotEsai      int    `gorm:"column:bobot_esai"`
	Opsi           int    `gorm:"column:opsi"`
	Date           string `gorm:"column:date"`
	Status         int    `gorm:"column:status"`
	SoalAgama      string `gorm:"column:soal_agama"`
	IdTp           int    `gorm:"column:id_tp"`
	IdSmt          int    `gorm:"column:id_smt"`
}
