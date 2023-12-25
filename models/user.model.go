package models

import (
	"fmt"
	"time"
)

type User struct {
	Id        int        `json:"id" gorm:"column:id;"`
	Name      string     `json:"name" gorm:"column:name;"`
	Age       int        `json:"age" gorm:"column:age;"`
	Address   string     `json:"address" gorm:"column:address;"`
	CreatedAt *time.Time `json:"createdAt" gorm:"column:created_at;"`
	UpdatedAt *time.Time `json:"updatedAt" gorm:"column:updated_at;"`
}

func (User) TableName() string {
	return "tbl_user"
}

func ToString(user User) string {
	return fmt.Sprintf("%d_%s", user.Id, user.Name)
}
