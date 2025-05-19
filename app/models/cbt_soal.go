package models

import "github.com/goravel/framework/database/orm"

type CbtSoal struct {
	orm.Model
	IdSoal      uint   `gorm:"column:id_soal;primaryKey"`
	BankId      int    `gorm:"column:bank_id"`
	MapelId     int    `gorm:"column:mapel_id"`
	Jenis       int    `gorm:"column:jenis"`
	NomorSoal   int    `gorm:"column:nomor_soal"`
	File        string `gorm:"column:file"`
	File1       string `gorm:"column:file1"`
	TipeFile    string `gorm:"column:tipe_file"`
	Soal        string `gorm:"column:soal"`
	OpsiA       string `gorm:"column:opsi_a"`
	OpsiB       string `gorm:"column:opsi_b"`
	OpsiC       string `gorm:"column:opsi_c"`
	OpsiD       string `gorm:"column:opsi_d"`
	OpsiE       string `gorm:"column:opsi_e"`
	FileA       string `gorm:"column:file_a"`
	FileB       string `gorm:"column:file_b"`
	FileC       string `gorm:"column:file_c"`
	FileD       string `gorm:"column:file_d"`
	FileE       string `gorm:"column:file_e"`
	Jawaban     string `gorm:"column:jawaban"`
	CreatedOn   int    `gorm:"column:created_on"`
	UpdatedOn   int    `gorm:"column:updated_on"`
}
