package models

import (
	"time"
)

type User struct {
	Id                  int
	Title               string
	Author              string
	ArticleCategoryId   int
	CoverMaterialBucket string
	Description         string
	Content             string
	ReadNum             int
	Status              int
	IsTop               int
	AdminId             int
	OrderTime           time.time
	CreatedTime         time.time
	UpdatedTime         time.Time
}

func (m *article) TableName() string {
	return TableName("article")
}
