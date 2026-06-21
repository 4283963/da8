package handlers

import (
	"jade-grading/models"
	"jade-grading/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type BraceletHandler struct {
	service     *services.BraceletService
	cardService *services.CardService
}

func NewBraceletHandler() *BraceletHandler {
	return &BraceletHandler{
		service:     services.NewBraceletService(),
		cardService: services.NewCardService(),
	}
}

type ApiResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func (h *BraceletHandler) CreateBracelet(c *gin.Context) {
	var req models.CreateBraceletRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, ApiResponse{
			Code:    400,
			Message: "参数错误: " + err.Error(),
			Data:    nil,
		})
		return
	}

	bracelet, err := h.service.CreateBracelet(&req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ApiResponse{
			Code:    500,
			Message: "创建失败: " + err.Error(),
			Data:    nil,
		})
		return
	}

	c.JSON(http.StatusOK, ApiResponse{
		Code:    200,
		Message: "创建成功",
		Data:    bracelet,
	})
}

func (h *BraceletHandler) GetAllBracelets(c *gin.Context) {
	bracelets, err := h.service.GetAllBracelets()
	if err != nil {
		c.JSON(http.StatusInternalServerError, ApiResponse{
			Code:    500,
			Message: "查询失败: " + err.Error(),
			Data:    nil,
		})
		return
	}

	c.JSON(http.StatusOK, ApiResponse{
		Code:    200,
		Message: "查询成功",
		Data:    bracelets,
	})
}

func (h *BraceletHandler) GetBracelet(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, ApiResponse{
			Code:    400,
			Message: "无效的ID",
			Data:    nil,
		})
		return
	}

	bracelet, err := h.service.GetBraceletByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, ApiResponse{
			Code:    404,
			Message: "记录不存在",
			Data:    nil,
		})
		return
	}

	c.JSON(http.StatusOK, ApiResponse{
		Code:    200,
		Message: "查询成功",
		Data:    bracelet,
	})
}

func (h *BraceletHandler) DeleteBracelet(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, ApiResponse{
			Code:    400,
			Message: "无效的ID",
			Data:    nil,
		})
		return
	}

	if err := h.service.DeleteBracelet(id); err != nil {
		c.JSON(http.StatusInternalServerError, ApiResponse{
			Code:    500,
			Message: "删除失败: " + err.Error(),
			Data:    nil,
		})
		return
	}

	c.JSON(http.StatusOK, ApiResponse{
		Code:    200,
		Message: "删除成功",
		Data:    nil,
	})
}

func (h *BraceletHandler) DownloadCard(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, ApiResponse{
			Code:    400,
			Message: "无效的ID",
			Data:    nil,
		})
		return
	}

	bracelet, err := h.service.GetBraceletByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, ApiResponse{
			Code:    404,
			Message: "记录不存在",
			Data:    nil,
		})
		return
	}

	cardContent, fileName := h.cardService.GeneratePrintCard(bracelet)

	c.Header("Content-Type", "text/plain; charset=utf-8")
	c.Header("Content-Disposition", "attachment; filename=\""+fileName+"\"")
	c.Header("Content-Transfer-Encoding", "binary")

	c.String(http.StatusOK, cardContent)
}
