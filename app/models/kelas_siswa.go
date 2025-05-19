package models

type KelasSiswa struct {
	IDKelasSiswa int `gorm:"column:id_kelas_siswa;primaryKey"`
	IDTp         int `gorm:"column:id_tp"`
	IDSmt        int `gorm:"column:id_smt"`
	IDSiswa      int `gorm:"column:id_siswa"`
	IDKelas      int `gorm:"column:id_kelas"`
}
