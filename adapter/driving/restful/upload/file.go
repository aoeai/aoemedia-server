package upload

import (
	"github.com/aoemedia-server/adapter/driving/restful/response"
	"github.com/aoemedia-server/application/image"
	"github.com/aoemedia-server/domain/file"
	"github.com/gin-gonic/gin"
)

// FileController 文件上传控制器
type FileController struct {
	BaseController
}

// NewFileController 创建文件上传控制器
func NewFileController() *FileController {
	return &FileController{}
}

// Upload 处理文件上传请求，接收并存储上传的文件，返回文件的基本信息
func (c *FileController) Upload(ctx *gin.Context) {
	content, originalFileName, err := c.readUploadedFile(ctx)
	if err != nil {
		return
	}

	fileContent := file.NewFileContent(content)
	service, err := image.NewFileStorage(fileContent)
	if err != nil {
		response.SendInternalServerError(ctx, "创建文件存储服务失败")
		return
	}

	_, saveErr := service.Save(originalFileName)
	if saveErr != nil {
		response.SendInternalServerError(ctx, saveErr.Error())
		return
	}

	c.sendSuccessResponse(ctx, 0, originalFileName, fileContent.SizeInBytes(), fileContent.Hash())
}
