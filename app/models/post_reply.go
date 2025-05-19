package models

import "github.com/goravel/framework/database/orm"

type PostReply struct {
	orm.Model
	IdReply    uint   `gorm:"column:id_reply;primaryKey"`
	IdComment  int    `gorm:"column:id_comment"`
	Dari       int    `gorm:"column:dari"`
	DariGroup  int    `gorm:"column:dari_group"`
	Text       string `gorm:"column:text"`
	Tanggal    string `gorm:"column:tanggal"`
	Updated    string `gorm:"column:updated"`
	Type       string `gorm:"column:type"`
}
