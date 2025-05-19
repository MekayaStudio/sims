package models

type KelasJadwalMapel struct {
	IDJadwal   int `gorm:"column:id_jadwal;primaryKey"`
	IDTp       int `gorm:"column:id_tp"`
	IDSmt      int `gorm:"column:id_smt"`
	IDKelas    int `gorm:"column:id_kelas"`
	IDHari     int `gorm:"column:id_hari"`
	JamKe      int `gorm:"column:jam_ke"`
	IDMapel    int `gorm:"column:id_mapel"`
}
