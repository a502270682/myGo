package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
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
	Name string
	Age  int
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
