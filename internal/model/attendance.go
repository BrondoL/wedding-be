package model

import "time"

type Attendance struct {
	ID          uint `gorm:"primaryKey"`
	Name        string
	Description string
	Status      string
	Number      *int
	CreatedAt   time.Time
}

func (Attendance) TableName() string {
	return "attendances"
}
