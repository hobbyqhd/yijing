package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hobbyqhd/yijing/service/models"
	"github.com/hobbyqhd/yijing/service/services"
)

type DivinationHandler struct {
	divinationService *services.DivinationService
}

func NewDivinationHandler() *DivinationHandler {
	return &DivinationHandler{
		divinationService: services.NewDivinationService(),
	}
}

func (h *DivinationHandler) CreateDivination(c *gin.Context) {
	var req services.DivinationRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的请求参数"})
		return
	}

	// 验证占卜类型
	if !models.IsValidDivinationType(req.Type) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "不支持的占卜类型"})
		return
	}

	// 获取用户ID
	userId := c.GetUint("userId")

	// 创建占卜记录
	divination, err := h.divinationService.CreateDivination(userId, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, divination)
}

func (h *DivinationHandler) GetUserDivinations(c *gin.Context) {
	userId := c.GetUint("userId")

	// 获取用户的占卜历史记录
	divinations, err := h.divinationService.GetUserDivinations(userId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, divinations)
}
