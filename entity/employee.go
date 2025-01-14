package entity

import	(
	"gorm.io/gorm"
)

type Employee struct {
	gorm.Model
	Name string `valid:"required~Name is required"`
	Email string `valid:"required~Email is required,email"`

	Room []Room `gorm:"foreignKey: employee_id"`
}