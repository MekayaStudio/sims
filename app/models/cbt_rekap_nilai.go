package models

import "github.com/goravel/framework/database/orm"

type CbtRekapNilai struct {
	orm.Model
	IdRekapNilai uint   `gorm:"column:id_rekap_nilai;primaryKey"`
	IdJadwal     int    `gorm:"column:id_jadwal"`
	IdTp         int    `gorm:"column:id_tp"`
	Tp           string `gorm:"column:tp"`
	IdSmt        int    `gorm:"column:id_smt"`
	Smt          string `gorm:"column:smt"`
	IdJenis      int    `gorm:"column:id_jenis"`
	KodeJenis    string `gorm:"column:kode_jenis"`
	IdBank       int    `gorm:"column:id_bank"`
	IdMapel      int    `gorm:"column:id_mapel"`
	IdSiswa      int    `gorm:"column:id_siswa"`
	IdKelas      int    `gorm:"column:id_kelas"`
	Kelas        string `gorm:"column:kelas"`
	Mulai        string `gorm:"column:mulai"`
	Selesai      string `gorm:"column:selesai"`
	Durasi       string `gorm:"column:durasi"`
	BobotPg      int    `gorm:"column:bobot_pg"`
	JawabanPg    string `gorm:"column:jawaban_pg"`
	NilaiPg      string `gorm:"column:nilai_pg"`
	BobotEsai    int    `gorm:"column:bobot_esai"`
	JawabanEsai  string `gorm:"column:jawaban_esai"`
	NilaiEsai    string `gorm:"column:nilai_esai"`
}
