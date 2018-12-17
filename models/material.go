package models

import (
	"time"
)

type User struct {
	Id          int
	Name        string
	Bucket      string
	AdminId     int
	CreatedTime time.time
	UpdatedTime time.time
}

func (m *material) TableName() string {
	return TableName("materials")
}
