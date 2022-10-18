package models

import (
	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type Users struct {
	gorm.Model
	Username string `gorm:"not null;uniqueIndex" json:"username" form:"username" valid:"required~Please input your username"`
	Email    string `gorm:"not null;uniqueIndex" json:"email" form:"email" valid:"required~Please input your email,email~Oops.. Your email is invalid. Please input your valid email"`
	Password string `gorm:"not null" json:"password" form:"password" valid:"required~Please input your password,minstringlength(6)~Oops.. Password has to have a minimum length of 6 characters"`
	Age      int    `gorm:"not null" json:"age" form:"age" valid:"required~Please input your age,range(8|70)~Oops.. your age not enough. Minimum age must be 8 years old"`
}

func (u *Users) BeforeCreate(tx *gorm.DB) (err error) {
	if _, err = govalidator.ValidateStruct(u); err != nil {
		return
	}
	return
}
func (u *Users) BeforeUpdate(tx *gorm.DB) (err error) {
	if _, err = govalidator.ValidateStruct(u); err != nil {
		return
	}
	return
}
