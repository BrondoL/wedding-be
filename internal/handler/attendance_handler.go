package handler

import (
	"net/http"

	"github.com/BrondoL/wedding-be/internal/dto/request"
	"github.com/BrondoL/wedding-be/internal/dto/response"
	"github.com/BrondoL/wedding-be/internal/util"
	"github.com/BrondoL/wedding-be/pkg/validator"
	"github.com/gin-gonic/gin"
)

func (h *Handler) GetAttendances(c *gin.Context) {
	ctx := c.Request.Context()
	attendances, err := h.attendanceService.GetAttendances(ctx)
	if err != nil {
		httpres := util.ResponseError(c, err, h.logger)
		c.JSON(httpres.Code, httpres)
		return
	}

	var res response.AttendancesResponse

	res.FormatAttendanceResponse(attendances)

	httpres := util.ResponseSuccess(c, h.logger, res, "get attendances success")
	c.JSON(httpres.Code, httpres)
}

func (h *Handler) CreateAttendance(c *gin.Context) {
	ctx := c.Request.Context()

	var req request.AttendanceRequest
	var res response.AttendanceResponse

	err := c.ShouldBindJSON(&req)
	if err != nil {
		validation := validator.FormatValidation(err)
		httpres := util.ResponseError(c, util.NewUnprocessibleEntityError(validation), h.logger)
		c.JSON(httpres.Code, httpres)
		return
	}

	m := req.ConvertToModel()
	attendance, err := h.attendanceService.CreateAttendance(ctx, m)
	if err != nil {
		httpres := util.ResponseError(c, err, h.logger)
		c.JSON(httpres.Code, httpres)
		return
	}

	res.FormatAttendance(attendance)

	httpres := util.ResponseSuccess(c, h.logger, res, "create attendance success", http.StatusCreated)
	c.JSON(httpres.Code, httpres)
}
