package models

import "github.com/goravel/framework/database/orm"

type UsersGroup struct {
	orm.Model
	ID      uint `gorm:"column:id;primaryKey"`
	UserID  uint `gorm:"column:user_id"`
	GroupID uint `gorm:"column:group_id"`
}
