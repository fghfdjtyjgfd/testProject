package service

import (
	"net/http"

	"github.com/gin-gonic/gin"
	model "github.com/jarbza/backend-api/repository/model"

	"github.com/jarbza/errs"
	// "github.com/jarbza/logx"
)

type Handler struct {
	service Servicer
}

func NewHandler(service Servicer) *Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) Login(c *gin.Context) {

	var req model.LoginRequest
	device, platform := c.GetHeader("x-device-model"), c.GetHeader("x-platform")

	if err := c.ShouldBindJSON(&req); err != nil {
		errs.JSON(c, errs.NewBadRequest(err.Error(), "Data require more"))
	}
	req.Device, req.Platform, req.ClientIP = device, platform, c.Copy().ClientIP()

	res, err := h.service.ValidateLogin(c, req)
	if err != nil {
		errs.JSON(c, err)
		return
	}

	c.JSON(http.StatusOK, res)
}

func (h *Handler) GetLoginHistory(c *gin.Context) {
	userRefID := c.GetHeader("userRefId")

	res, err := h.service.GetRecords(c, userRefID)
	if err != nil {
		errs.JSON(c, err)
		return
	}

	c.JSON(http.StatusOK, res)
}

func (h *Handler) GetStatus(c *gin.Context) {
	var req model.LoginHistory
	req.UserRefID = c.GetHeader("userRefId")
	req.Device, req.Platform = c.GetHeader("x-device-model"), c.GetHeader("x-platform")

	res := h.service.CheckStatus(c, req)
	c.JSON(http.StatusOK, res)
}
