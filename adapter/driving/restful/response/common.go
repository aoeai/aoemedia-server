package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// SendSuccess 发送成功响应
func SendSuccess(ctx *gin.Context, data gin.H) {
	ctx.JSON(http.StatusOK, data)
}

// SendBadRequest 请求参数错误
func SendBadRequest(ctx *gin.Context, errorMsg string) {
	sendError(ctx, http.StatusBadRequest, errorMsg)
}

// SendInternalServerError 服务器内部错误
func SendInternalServerError(ctx *gin.Context, errorMsg string) {
	sendError(ctx, http.StatusInternalServerError, errorMsg)
}

// sendError 统一错误响应处理
func sendError(ctx *gin.Context, statusCode int, errorMsg string) {
	ctx.JSON(statusCode, gin.H{"error": errorMsg})
}

func SendUnauthorized(ctx *gin.Context) {
	sendError(ctx, http.StatusUnauthorized, "无效的认证")
}
