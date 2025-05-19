package models

import "github.com/goravel/framework/database/orm"

type Post struct {
	orm.Model
	IdPost     uint   `gorm:"column:id_post;primaryKey"`
	Dari       int    `gorm:"column:dari"`
	DariGroup  int    `gorm:"column:dari_group"`
	Kepada     string `gorm:"column:kepada"`
	Text       string `gorm:"column:text"`
	Tanggal    string `gorm:"column:tanggal"`
	Updated    string `gorm:"column:updated"`
}
