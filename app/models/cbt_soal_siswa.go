package models

import "github.com/goravel/framework/database/orm"

type CbtSoalSiswa struct {
	orm.Model
	IdSoalSiswa   string `gorm:"column:id_soal_siswa;primaryKey"`
	IdBank        int    `gorm:"column:id_bank"`
	IdJadwal      int    `gorm:"column:id_jadwal"`
	IdSoal        int    `gorm:"column:id_soal"`
	IdSiswa       int    `gorm:"column:id_siswa"`
	JenisSoal     int    `gorm:"column:jenis_soal"`
	NoSoalAlias   int    `gorm:"column:no_soal_alias"`
	OpsiAliasA    string `gorm:"column:opsi_alias_a"`
	OpsiAliasB    string `gorm:"column:opsi_alias_b"`
	OpsiAliasC    string `gorm:"column:opsi_alias_c"`
	OpsiAliasD    string `gorm:"column:opsi_alias_d"`
	OpsiAliasE    string `gorm:"column:opsi_alias_e"`
	JawabanAlias  string `gorm:"column:jawaban_alias"`
	JawabanSiswa  string `gorm:"column:jawaban_siswa"`
	JawabanBenar  string `gorm:"column:jawaban_benar"`
	PointEssai    int    `gorm:"column:point_essai"`
	SoalEnd       int    `gorm:"column:soal_end"`
	PointSoal     string `gorm:"column:point_soal"`
	NilaiKoreksi  string `gorm:"column:nilai_koreksi"`
	NilaiOtomatis int    `gorm:"column:nilai_otomatis"`
	TimeCreate    string `gorm:"column:time_create"`
}
