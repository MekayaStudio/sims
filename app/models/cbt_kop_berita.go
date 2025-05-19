package models

import "github.com/goravel/framework/database/orm"

type CbtKopBerita struct {
	orm.Model
	Header1 string `gorm:"column:header_1"`
	Header2 string `gorm:"column:header_2"`
	Header3 string `gorm:"column:header_3"`
	Header4 string `gorm:"column:header_4"`
}
