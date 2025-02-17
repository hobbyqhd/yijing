package handlers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/hobbyqhd/yijing/service/services"
)

type FortuneHandler struct {
	fortuneService *services.FortuneService
}

func NewFortuneHandler() *FortuneHandler {
	return &FortuneHandler{
		fortuneService: services.NewFortuneService(),
	}
}

func (h *FortuneHandler) CalculateFortune(c *gin.Context) {
	userId := c.GetUint("userId")

	fortune, err := h.fortuneService.CalculateFortune(userId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, fortune)
}

type GetFortunesRequest struct {
	StartDate string `form:"start_date" binding:"required"`
	EndDate   string `form:"end_date" binding:"required"`
}

func (h *FortuneHandler) GetUserFortunes(c *gin.Context) {
	var req GetFortunesRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的请求参数"})
		return
	}

	startDate, err := time.Parse("2006-01-02", req.StartDate)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的开始日期格式"})
		return
	}

	endDate, err := time.Parse("2006-01-02", req.EndDate)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的结束日期格式"})
		return
	}

	userId := c.GetUint("userId")
	fortunes, err := h.fortuneService.GetUserFortunes(userId, startDate, endDate)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, fortunes)
}
