package handler

import (
	"github.com/BrondoL/wedding-be/internal/util"
	"github.com/gin-gonic/gin"
)

func (h *Handler) NoRoute(c *gin.Context) {
	httpres := util.ResponseError(c, util.NewNotFoundError("page not found"), h.logger)
	c.JSON(httpres.Code, httpres)
}
