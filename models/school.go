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

var schoolDao *SchoolDao

type SchoolDao struct {
	DB *gorm.DB
}

func InitSchoolDao(d *gorm.DB) {
	schoolDao = &SchoolDao{
		DB: d,
	}
}

func GetSchoolDao() *SchoolDao {
	return schoolDao
}

func (s *SchoolDao) GetSchoolsByNames(names []string) ([]*School, error) {
	var ret []*School
	if err := s.DB.Where("name in (?)", names).Find(&ret).Error; err != nil {
		return nil, err
	}
	return ret, nil
}
