package models

import "github.com/goravel/framework/database/orm"

type ApiSetting struct {
	orm.Model
	ID                uint `gorm:"column:id;primaryKey"`
	AutoSync          int  `gorm:"column:auto_sync"`
	EditProfileSiswa  int  `gorm:"column:edit_profile_siswa"`
	EditProfileGuru   int  `gorm:"column:edit_profile_guru"`
}
