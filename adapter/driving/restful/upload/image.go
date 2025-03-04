package upload

import (
	"fmt"
	filerepo "github.com/aoemedia-server/adapter/driven/persistence/file"
	"github.com/aoemedia-server/adapter/driving/restful/response"
	"github.com/aoemedia-server/application/image"
	"github.com/aoemedia-server/domain/file"
	imagemodel "github.com/aoemedia-server/domain/image"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"strconv"
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
		return // readUploadedFile内部已处理错误响应
	}

	// 获取 source 参数的值
	source, err := parseSource(ctx)
	if err != nil {
		response.SendBadRequest(ctx, err.Error())
		return
	}
	logrus.Infof("source: %d", source)

	fileContent := file.NewFileContent(content)
	domainImage, err := imagemodel.NewDomainImage(fileContent)
	if err != nil {
		response.SendBadRequest(ctx, err.Error())
		return
	}

	storage, err := image.NewImageStorage(domainImage, filerepo.NewRepository())
	if err != nil {
		logrus.Error("创建图片服务失败: ", err)
		response.SendInternalServerError(ctx, "创建图片服务失败")
		return
	}

	id, save, err := storage.Save(originalFileName)
	if err != nil {
		logrus.Error("保存图片失败: ", err)
		response.SendInternalServerError(ctx, "保存图片失败")
		return
	}
	logrus.Infof("图片保存成功: %s", save)

	c.sendSuccessResponse(ctx, id, originalFileName, fileContent.SizeInBytes(), fileContent.Hash())
}

// parseSource 从请求中解析source参数
func parseSource(ctx *gin.Context) (uint8, error) {
	source := uint8(1) // 默认值
	if sourceStr := ctx.PostForm("source"); sourceStr != "" {
		sourceInt, err := strconv.ParseUint(sourceStr, 10, 8)
		if err != nil {
			return 0, fmt.Errorf("source 参数值无效")
		}
		source = uint8(sourceInt)
	}
	return source, nil
}
