package upload

import (
	"github.com/aoemedia-server/adapter/driving/restful/response"
	"github.com/aoemedia-server/common/converter"
	"github.com/gin-gonic/gin"
	"mime/multipart"
)

// BaseController 基础控制器，包含文件上传的通用功能
type BaseController struct{}

// readUploadedFile 读取上传的文件内容
func (c *BaseController) readUploadedFile(ctx *gin.Context) ([]byte, string, error) {
	// 获取上传的文件
	file, header, err := ctx.Request.FormFile("file")
	if err != nil {
		response.SendBadRequest(ctx, "未找到上传文件")
		return nil, "", err
	}
	defer func(file multipart.File) {
		err := file.Close()
		if err != nil {
			response.SendInternalServerError(ctx, "关闭文件失败")
		}
	}(file)

	// 读取文件内容
	content := make([]byte, header.Size)
	if _, err := file.Read(content); err != nil {
		response.SendInternalServerError(ctx, "读取文件失败")
		return nil, "", err
	}

	return content, header.Filename, nil
}

// sendSuccessResponse 发送成功响应
func (c *BaseController) sendSuccessResponse(ctx *gin.Context, id int64, filename string, size int64, hash string) {
	response.SendSuccess(ctx, gin.H{
		"message":  "文件上传成功",
		"id":       converter.Int64ToStr(id),
		"filename": filename,
		"size":     size,
		"hash":     hash,
	})
}
