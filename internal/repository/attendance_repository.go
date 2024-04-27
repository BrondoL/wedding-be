package repository

import (
	"github.com/BrondoL/wedding-be/internal/model"
	"gorm.io/gorm"
)

type AttendanceRepository interface {
	FindAll() ([]*model.Attendance, error)
	Save(*model.Attendance) (*model.Attendance, error)
}

type attendanceRepository struct {
	db *gorm.DB
}

type ARConfig struct {
	DB *gorm.DB
}

func NewAttendanceRepository(c *ARConfig) AttendanceRepository {
	return &attendanceRepository{
		db: c.DB,
	}
}

func (r *attendanceRepository) FindAll() ([]*model.Attendance, error) {
	var attendances []*model.Attendance

	err := r.db.Order("created_at DESC").Find(&attendances).Error
	if err != nil {
		return attendances, err
	}

	return attendances, nil
}

func (r *attendanceRepository) Save(m *model.Attendance) (*model.Attendance, error) {
	err := r.db.Create(m).Error
	if err != nil {
		return nil, err
	}

	return m, nil
}
