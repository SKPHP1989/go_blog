package models

import (
	"time"
)

type User struct {
	Id        int
	Username  string
	Password  string
	Mask      string
	Email     string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (m *User) TableName() string {
	return TableName("user")
}
