package models

import (
	"golang.org/x/crypto/bcrypt"
	"time"
)

type User struct {
	Id        int        `json:"id" gorm:"column:id;"`
	Name      string     `json:"name" gorm:"column:name;"`
	Age       int        `json:"age" gorm:"column:age;"`
	Address   string     `json:"address" gorm:"column:address;"`
	Email     string     `json:"email" gorm:"column:email"`
	Password  string     `json:"password" gorm:"column:password"`
	RoleId    int        `json:"roleId" gorm:"column:role_id"`
	CreatedAt *time.Time `json:"createdAt" gorm:"column:created_at;"`
	UpdatedAt *time.Time `json:"updatedAt" gorm:"column:updated_at;"`
}

func (user *User) TableName() string {
	return "tbl_user"
}

func (user *User) HashPassword(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user.Password = string(bytes)
	return nil
}

func (user *User) CheckMatchesPassword(password string) error {
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return err
	}

	return nil
}
