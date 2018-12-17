package models

import (
	"time"
)

type User struct {
	Id                  int
	Name                string
	CoverMaterialBucket string
	Description         string
	Status              int
	is_top              int
	OrderTime           time.time
	CreatedTime         time.time
	UpdatedTime         time.time
}

func (m *articleCategory) TableName() string {
	return TableName("article_categories")
}
