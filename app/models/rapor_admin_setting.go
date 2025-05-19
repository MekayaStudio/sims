package models

import "github.com/goravel/framework/database/orm"

type RaporAdminSetting struct {
	orm.Model
	IdSetting           uint   `gorm:"column:id_setting;primaryKey"`
	IdTp                int    `gorm:"column:id_tp"`
	IdSmt               int    `gorm:"column:id_smt"`
	KkmTunggal          int    `gorm:"column:kkm_tunggal"`
	Kkm                 int    `gorm:"column:kkm"`
	BobotPh             int    `gorm:"column:bobot_ph"`
	BobotPts            int    `gorm:"column:bobot_pts"`
	BobotPas            int    `gorm:"column:bobot_pas"`
	BobotAbsen          int    `gorm:"column:bobot_absen"`
	TglRaporAkhir       string `gorm:"column:tgl_rapor_akhir"`
	TglRaporKelasAkhir  string `gorm:"column:tgl_rapor_kelas_akhir"`
	TglRaporPts         string `gorm:"column:tgl_rapor_pts"`
	NipKepsek           int    `gorm:"column:nip_kepsek"`
	NipWalikelas        int    `gorm:"column:nip_walikelas"`
}
