package models

import "github.com/goravel/framework/database/orm"

type Setting struct {
	orm.Model
	IdSetting     uint   `gorm:"column:id_setting;primaryKey"`
	KodeSekolah   string `gorm:"column:kode_sekolah"`
	Sekolah       string `gorm:"column:sekolah"`
	Npsn          string `gorm:"column:npsn"`
	Nss           string `gorm:"column:nss"`
	Jenjang       int    `gorm:"column:jenjang"`
	Kepsek        string `gorm:"column:kepsek"`
	Nip           string `gorm:"column:nip"`
	TandaTangan   string `gorm:"column:tanda_tangan"`
	Alamat        string `gorm:"column:alamat"`
	Desa          string `gorm:"column:desa"`
	Kecamatan     string `gorm:"column:kecamatan"`
	Kota          string `gorm:"column:kota"`
	Provinsi      string `gorm:"column:provinsi"`
	KodePos       int    `gorm:"column:kode_pos"`
	Telp          string `gorm:"column:telp"`
	Fax           string `gorm:"column:fax"`
	Web           string `gorm:"column:web"`
	Email         string `gorm:"column:email"`
	NamaAplikasi  string `gorm:"column:nama_aplikasi"`
	LogoKanan     string `gorm:"column:logo_kanan"`
	LogoKiri      string `gorm:"column:logo_kiri"`
}
