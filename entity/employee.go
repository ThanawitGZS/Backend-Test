package entity

import	(
	"gorm.io/gorm"
)

type Employee struct {
	gorm.Model
	Name string `valid:"required~Name is required"`

	Room []Room `gorm:"foreignKey: employee_id"`
}