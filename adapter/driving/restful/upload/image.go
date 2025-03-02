package upload

import (
	"github.com/aoemedia-server/adapter/driving/restful/response"
	"github.com/aoemedia-server/application/storage"
	"github.com/aoemedia-server/domain/file"
	imagemodel "github.com/aoemedia-server/domain/image"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// ImageController 图片上传控制器
type ImageController struct {
	BaseController
}

// NewImageController 创建图片上传控制器
func NewImageController() *ImageController {
	return &ImageController{}
}

// Upload 处理图片上传请求
func (c *ImageController) Upload(ctx *gin.Context) {
	content, originalFileName, err := c.readUploadedFile(ctx)
	if err != nil {
		return // 假设readUploadedFile内部已处理错误响应
	}

	fileContent := file.NewFileContent(content)
	aoeImage, err := imagemodel.NewAoeImage(fileContent)
	if err != nil {
		response.SendBadRequest(ctx, err.Error())
		return
	}

	service, err := storage.NewImageStorage(aoeImage)
	if err != nil {
		response.SendInternalServerError(ctx, "创建图片服务失败")
		return
	}

	save, err := service.Save(originalFileName)
	if err != nil {
		response.SendInternalServerError(ctx, "保存图片失败")
		return
	}
	logrus.Infof("图片保存成功: %s", save)

	c.sendSuccessResponse(ctx, originalFileName, fileContent.SizeInBytes(), fileContent.Hash())
}
