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

func FormatAttendanceResponse(attendances []*model.Attendance) AttendancesResponse {
	total, hadir, tidakHadir, ragu := 0, 0, 0, 0

	formattedAttendances := []AttendanceResponse{}

	for _, attendance := range attendances {
		if attendance.Status == constant.StatusHadir {
			if attendance.Number != nil {
				hadir += *attendance.Number
				total += *attendance.Number
			}
		} else if attendance.Status == constant.StatusTidakHadir {
			tidakHadir += 1
			total += 1
		} else {
			ragu += 1
			total += 1
		}

		formattedAttendance := FormatAttendance(attendance)
		formattedAttendances = append(formattedAttendances, formattedAttendance)
	}

	formattedSummary := Summary{}
	formattedSummary.Total = total
	formattedSummary.Hadir = hadir
	formattedSummary.TidakHadir = tidakHadir
	formattedSummary.Ragu = ragu

	return AttendancesResponse{
		Attendances: formattedAttendances,
		Summary:     formattedSummary,
	}
}

func FormatAttendance(attendance *model.Attendance) AttendanceResponse {
	formattedAttendance := AttendanceResponse{}
	formattedAttendance.Name = attendance.Name
	formattedAttendance.Description = attendance.Description
	formattedAttendance.Status = attendance.Status
	formattedAttendance.Number = attendance.Number
	formattedAttendance.CreatedAt = attendance.CreatedAt.Format("2006-01-02 15:04:05")

	return formattedAttendance
}
