package models

import (
	"time"
)

type User struct {
	Id           int
	Username     string
	Password     string
	PasswordMask string
	Email        string
	LastLoginIp  string
	CreatedTime  time.Time
	UpdatedTime  time.Time
}

func (m *admin) TableName() string {
	return TableName("admins")
}
