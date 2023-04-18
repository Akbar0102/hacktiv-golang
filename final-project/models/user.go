package models

import (
	"final-project/helpers"

	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

func init() {
	govalidator.CustomTypeTagMap.Set("checkAgeValidator", govalidator.CustomTypeValidator(func(i interface{}, context interface{}) bool {
		if age, ok := i.(int); ok {
			if age > 8 {
				return true
			}
		}
		return false
	}))
}

type User struct {
	GormModel
	Username     string        `gorm:"not null;uniqueIndex;column:username" json:"username" valid:"required~Your username is required"`
	Email        string        `gorm:"not null;uniqueIndex;column:email" json:"email" valid:"required~Your email is required,email~Invalid email format"`
	Password     string        `gorm:"not null;column:password" json:"password" valid:"required~Your password is required,minstringlength(6)~Password has to have a minimum length of 6 characters"`
	Age          int           `gorm:"not null;column:age" json:"age" valid:"required~Your age is required,checkAgeValidator~Age must be greater than 8"`
	Photos       []Photo       `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"photos"`
	Comments     []Comment     `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"comments"`
	SocialMedias []SocialMedia `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"social_medias"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(u)

	if errCreate != nil {
		err = errCreate
		return
	}

	u.Password = helpers.HashPass(u.Password)

	err = nil
	return
}
