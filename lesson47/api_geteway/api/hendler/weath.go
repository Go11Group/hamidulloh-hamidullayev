package handler

import (
	"dodi/genproto/protos"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) Get(ctx *gin.Context) {
	req := &protos.CurrentWeatherRequest{}

	resp, err := h.Weather.GetCurrentWeather(ctx, req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, resp)
}