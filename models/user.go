package models

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	Mobile string `gorm:"type:char(11);index;unique;not null;" json:"mobile,omitempty"`
}

//根据手机号码查找用户
func FindUserByMobile(mobile string) (*User, error) {
	var user User
	err := db.Where("mobile=?", mobile).First(&user).Error
	return &user, err
}

//创建用户
func CreateUser(mobile string) (*User, error) {
	user := User{
		Mobile: mobile,
	}
	err := db.Create(&user).Error
	return &user, err
}
