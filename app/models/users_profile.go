package models

import "github.com/goravel/framework/database/orm"

type UsersProfile struct {
	orm.Model
	IdUser      uint   `gorm:"column:id_user;primaryKey"`
	NamaLengkap string `gorm:"column:nama_lengkap"`
	Jabatan     string `gorm:"column:jabatan"`
	LevelAccess int    `gorm:"column:level_access"`
	Foto        string `gorm:"column:foto"`
}
