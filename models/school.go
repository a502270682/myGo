package models

import "github.com/jinzhu/gorm"

const TableSchool = "school"

type School struct {
	*gorm.Model
	Name string
}

func (s *School) TableName() string {
	return TableSchool
}
