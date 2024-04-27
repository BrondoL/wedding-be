package request

import "github.com/BrondoL/wedding-be/internal/model"

type AttendanceRequest struct {
	Name        string `json:"name" binding:"required,min=5,max=20"`
	Description string `json:"description" binding:"required,min=5"`
	Status      string `json:"status" binding:"required"`
	Number      *int   `json:"number"`
}

func (r *AttendanceRequest) ConvertToModel() *model.Attendance {
	return &model.Attendance{
		Name:        r.Name,
		Description: r.Description,
		Status:      r.Status,
		Number:      r.Number,
	}
}
