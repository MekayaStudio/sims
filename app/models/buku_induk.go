package models

import "github.com/goravel/framework/database/orm"

type BukuInduk struct {
	orm.Model
	IdSiswa           uint   `gorm:"column:id_siswa;primaryKey"`
	Uid               string `gorm:"column:uid"`
	RombelAwal        string `gorm:"column:rombel_awal"`
	NamaPanggilan     string `gorm:"column:nama_panggilan"`
	Bahasa            string `gorm:"column:bahasa"`
	JmlSaudaraKandung int    `gorm:"column:jml_saudara_kandung"`
	JmlSaudaraTiri    int    `gorm:"column:jml_saudara_tiri"`
	JmlSaudaraAngkat  int    `gorm:"column:jml_saudara_angkat"`
	Yatim             int    `gorm:"column:yatim"`
	TinggalBersama    string `gorm:"column:tinggal_bersama"`
	Jarak             string `gorm:"column:jarak"`
	GolDarah          string `gorm:"column:gol_darah"`
	Penyakit          string `gorm:"column:penyakit"`
	KelainanFisik     string `gorm:"column:kelainan_fisik"`
	Kegemaran         string `gorm:"column:kegemaran"`
	Beasiswa          string `gorm:"column:beasiswa"`
	NoIjazahSebelumnya string `gorm:"column:no_ijazah_sebelumnya"`
	TahunLulusSebelumnya string `gorm:"column:tahun_lulus_sebelumnya"`
	PindahanDari      string `gorm:"column:pindahan_dari"`
	AlasanKepindahan  string `gorm:"column:alasan_kepindahan"`
	AgamaAyah         string `gorm:"column:agama_ayah"`
	TempatLahirAyah   string `gorm:"column:tempat_lahir_ayah"`
}
