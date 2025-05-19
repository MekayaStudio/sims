package models

import "github.com/goravel/framework/database/orm"

type MasterSiswa struct {
	orm.Model
	IdSiswa       uint   `gorm:"column:id_siswa;primaryKey"`
	Nisn          uint64 `gorm:"column:nisn"`
	Nis           string `gorm:"column:nis"`
	Nama          string `gorm:"column:nama"`
	JenisKelamin  string `gorm:"column:jenis_kelamin"`
	Username      string `gorm:"column:username"`
	Password      string `gorm:"column:password"`
	KelasAwal     int    `gorm:"column:kelas_awal"`
	TahunMasuk    string `gorm:"column:tahun_masuk"`
	SekolahAsal   string `gorm:"column:sekolah_asal"`
	TempatLahir   string `gorm:"column:tempat_lahir"`
	TanggalLahir  string `gorm:"column:tanggal_lahir"`
	Agama         string `gorm:"column:agama"`
	Hp            string `gorm:"column:hp"`
	Email         string `gorm:"column:email"`
	Foto          string `gorm:"column:foto"`
	AnakKe        int    `gorm:"column:anak_ke"`
	StatusKeluarga string `gorm:"column:status_keluarga"`
	Alamat        string `gorm:"column:alamat"`
	Rt            string `gorm:"column:rt"`
	Rw            string `gorm:"column:rw"`
	Kelurahan     string `gorm:"column:kelurahan"`
}
