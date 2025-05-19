package models

import "github.com/goravel/framework/database/orm"

type MasterSmt struct {
	orm.Model
	IdSmt    uint   `gorm:"column:id_smt;primaryKey"`
	Smt      string `gorm:"column:smt"`
	NamaSmt  string `gorm:"column:nama_smt"`
	Active   int    `gorm:"column:active"`
}
