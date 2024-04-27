package handler

import (
	"github.com/BrondoL/wedding-be/internal/util"
	"github.com/gin-gonic/gin"
)

func (h *Handler) DefaultHandler(c *gin.Context) {
	httpres := util.ResponseSuccess(nil, "Welcome to Levri & Nabil Wedding API")
	c.JSON(httpres.Code, httpres)
}
