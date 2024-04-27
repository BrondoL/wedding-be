package response

import (
	"github.com/BrondoL/wedding-be/internal/constant"
	"github.com/BrondoL/wedding-be/internal/model"
)

type AttendancesResponse struct {
	Attendances []AttendanceResponse `json:"attendances"`
	Summary     Summary              `json:"summary"`
}

type AttendanceResponse struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Status      string `json:"status"`
	Number      *int   `json:"number"`
	CreatedAt   string `json:"created_at"`
}

type Summary struct {
	Total      int `json:"total"`
	Hadir      int `json:"hadir"`
	TidakHadir int `json:"tidak_hadir"`
	Ragu       int `json:"ragu"`
}

func (r *AttendancesResponse) FormatAttendanceResponse(attendances []*model.Attendance) {
	total, hadir, tidakHadir, ragu := 0, 0, 0, 0

	formattedAttendances := []AttendanceResponse{}
	for _, attendance := range attendances {
		if attendance.Status == constant.StatusHadir {
			hadir += *attendance.Number
			total += *attendance.Number
		} else if attendance.Status == constant.StatusTidakHadir {
			tidakHadir += 1
			total += 1
		} else {
			ragu += 1
			total += 1
		}

		formattedAttendance := AttendanceResponse{
			Name:        attendance.Name,
			Description: attendance.Description,
			Status:      attendance.Status,
			Number:      attendance.Number,
			CreatedAt:   attendance.CreatedAt.Format("2006-01-02 15:04:05"),
		}
		formattedAttendances = append(formattedAttendances, formattedAttendance)
	}

	r.Attendances = formattedAttendances
	r.Summary.Total = total
	r.Summary.Hadir = hadir
	r.Summary.TidakHadir = tidakHadir
	r.Summary.Ragu = ragu
}

func (r *AttendanceResponse) FormatAttendance(attendance *model.Attendance) {
	r.Name = attendance.Name
	r.Description = attendance.Description
	r.Status = attendance.Status
	r.Number = attendance.Number
	r.CreatedAt = attendance.CreatedAt.Format("2006-01-02 15:04:05")
}
