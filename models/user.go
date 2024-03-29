package models

import (
	"fmt"

	"gorm.io/gorm"
)

const TableUser = "user"

var userDao *UserDao

type UserDao struct {
	DB *gorm.DB
}

func InitUserDao(d *gorm.DB) {
	userDao = &UserDao{
		DB: d,
	}
}

func GetUserDao() *UserDao {
	return userDao
}

type User struct {
	*gorm.Model
	Name     string `gorm:"column:name"`
	Age      int    `gorm:"column:age"`
	SchoolId uint   `gorm:"column:school_id"`
}

func (u *User) TableName() string {
	return TableUser
}

func (d *UserDao) CreatUser(u *User) error {
	return d.DB.Create(u).Error
}

func (d *UserDao) GetUser(name string) (*User, error) {
	user := &User{}
	err := d.DB.Where("name = (?)", name).First(user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (d *UserDao) UpdateMany() error {
	var err error
	var users []*User
	if err = d.DB.Where("age > (?)", 16).Find(&users).Error; err != nil {
		panic(err)
	}
	for i := range users {
		users[i].Age = 15
	}
	if err = d.DB.Where("age > (?)", 16).UpdateColumns(&users).Error; err != nil {
		panic(err)
	}
	var newUsers []*User
	if err = d.DB.Where("id = (?)", users[0].ID).Find(&newUsers).Error; err != nil {
		panic(err)
	}
	fmt.Println(newUsers[0])
	return nil
}

type Result struct {
	UserName   string `json:"user_name"`
	SchoolName string `json:"school_name"`
	Age        int    `json:"age"`
	SchoolId   uint   `json:"school_id"`
}

func (d *UserDao) GetUserWithSchool(name string) []*Result {
	// .Select("user.name as username, school.name as schoolname")
	var ret []*Result
	if err := d.DB.Table("user").Where("user.name = (?)", name).Select("user.name as user_name, school.name as school_name, user.age as age, user.school_id as school_id").Joins("left join school on user.school_id = school.id").Find(&ret).Error; err != nil {
		panic(err)
	}
	return ret
}
