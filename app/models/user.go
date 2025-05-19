package models

import (
	"github.com/goravel/framework/database/orm"
)

type User struct {
	orm.Model
	IpAddress                  string  `gorm:"column:ip_address"`
	Username                   *string `gorm:"column:username"`
	Password                   string  `gorm:"column:password"`
	Email                      *string `gorm:"column:email"`
	ActivationSelector         *string `gorm:"column:activation_selector"`
	ActivationCode             *string `gorm:"column:activation_code"`
	ForgottenPasswordSelector  *string `gorm:"column:forgotten_password_selector"`
	ForgottenPasswordCode      *string `gorm:"column:forgotten_password_code"`
	ForgottenPasswordTime      *uint   `gorm:"column:forgotten_password_time"`
	RememberSelector           *string `gorm:"column:remember_selector"`
	RememberCode               *string `gorm:"column:remember_code"`
	CreatedOn                  uint    `gorm:"column:created_on"`
	LastLogin                  *uint   `gorm:"column:last_login"`
	Active                     *uint8  `gorm:"column:active"`
	FirstName                  *string `gorm:"column:first_name"`
	LastName                   *string `gorm:"column:last_name"`
	Company                    *string `gorm:"column:company"`
	Phone                      *string `gorm:"column:phone"`
}
