package models

type KelasJadwalMateri struct {
	IDKJM      string `gorm:"column:id_kjm;primaryKey"`
	IDTp       int    `gorm:"column:id_tp"`
	IDSmt      int    `gorm:"column:id_smt"`
	IDMateri   int    `gorm:"column:id_materi"`
	IDMapel    int    `gorm:"column:id_mapel"`
	IDKelas    int    `gorm:"column:id_kelas"`
	JadwalMateri string `gorm:"column:jadwal_materi"`
	Jenis      int    `gorm:"column:jenis"`
}
