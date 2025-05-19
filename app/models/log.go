package models

type Log struct {
	IDLog     int    `gorm:"column:id_log;primaryKey"`
	LogTime   string `gorm:"column:log_time"`
	IDUser    int    `gorm:"column:id_user"`
	IDGroup   int    `gorm:"column:id_group"`
	NameGroup string `gorm:"column:name_group"`
	LogType   int    `gorm:"column:log_type"`
	LogDesc   string `gorm:"column:log_desc"`
	Address   string `gorm:"column:address"`
	Agent     string `gorm:"column:agent"`
	Device    string `gorm:"column:device"`
}
