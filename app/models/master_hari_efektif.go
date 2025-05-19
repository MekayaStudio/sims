package models

import "github.com/goravel/framework/database/orm"

type MasterHariEfektif struct {
	orm.Model
	IdHariEfektif   uint `gorm:"column:id_hari_efektif;primaryKey"`
	JmlHariEfektif  int  `gorm:"column:jml_hari_efektif"`
}
