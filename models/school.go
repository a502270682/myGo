package models

import "github.com/jinzhu/gorm"

const TableSchool = "school"

type SchoolType string

const SchoolTypeHighSchool SchoolType = "h school"
const SchoolTypeJSchool SchoolType = "j school"

type School struct {
	*gorm.Model
	Name       string     `gorm:"column:name"`
	SchoolType SchoolType `gorm:"column:school_type"`
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

func (s *SchoolDao) GetSchoolsByType(types []SchoolType) ([]*School, error) {
	var ret []*School
	if err := s.DB.Where("school_type in (?)", types).Find(&ret).Error; err != nil {
		return nil, err
	}
	return ret, nil
}
