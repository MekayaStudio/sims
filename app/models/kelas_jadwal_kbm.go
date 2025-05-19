package models

type KelasJadwalKbm struct {
	IDKBM         int    `gorm:"column:id_kbm;primaryKey"`
	IDTp          int    `gorm:"column:id_tp"`
	IDSmt         int    `gorm:"column:id_smt"`
	IDKelas       int    `gorm:"column:id_kelas"`
	KBMJamPel     int    `gorm:"column:kbm_jam_pel"`
	KBMJamMulai   string `gorm:"column:kbm_jam_mulai"`
	KBMJmlMapel   int    `gorm:"column:kbm_jml_mapel_hari"`
	Istirahat     string `gorm:"column:istirahat"`
}
