package models

import "github.com/goravel/framework/database/orm"

type PostComment struct {
	orm.Model
	IdComment  uint   `gorm:"column:id_comment;primaryKey"`
	IdPost     int    `gorm:"column:id_post"`
	Dari       int    `gorm:"column:dari"`
	DariGroup  int    `gorm:"column:dari_group"`
	Text       string `gorm:"column:text"`
	Tanggal    string `gorm:"column:tanggal"`
	Updated    string `gorm:"column:updated"`
	Type       string `gorm:"column:type"`
}
