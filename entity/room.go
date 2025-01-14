package entity

import (
	"gorm.io/gorm"
)

type Room struct {
	gorm.Model
	RoomName string `valid:"required~RoomName is required"`

	EmployeeID uint `valid:"required~EmployeeID is required"`
	Employee Employee `gorm:"foreignKey: employee_id"`
}