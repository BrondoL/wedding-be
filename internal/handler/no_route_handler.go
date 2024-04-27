package handler

import (
	"github.com/BrondoL/wedding-be/internal/util"
	"github.com/gin-gonic/gin"
)

func (h *Handler) NoRoute(c *gin.Context) {
	httpres := util.ResponseError(util.NewNotFoundError("page not found"))
	c.JSON(httpres.Code, httpres)
}
