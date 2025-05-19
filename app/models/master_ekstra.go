package models

import "github.com/goravel/framework/database/orm"

type MasterEkstra struct {
	orm.Model
	IdEkstra   uint   `gorm:"column:id_ekstra;primaryKey"`
	NamaEkstra string `gorm:"column:nama_ekstra"`
	KodeEkstra string `gorm:"column:kode_ekstra"`
}
