package models

type JabatanGuru struct {
	IDJabatanGuru string `gorm:"column:id_jabatan_guru;primaryKey"`
	IDGuru        int    `gorm:"column:id_guru"`
	IDJabatan     int    `gorm:"column:id_jabatan"`
	IDKelas       int    `gorm:"column:id_kelas"`
	MapelKelas    string `gorm:"column:mapel_kelas"`
	EkstraKelas   string `gorm:"column:ekstra_kelas"`
	IDTp          int    `gorm:"column:id_tp"`
	IDSmt         int    `gorm:"column:id_smt"`
}
