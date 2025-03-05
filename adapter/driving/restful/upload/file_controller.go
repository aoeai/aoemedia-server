package upload

import (
	"github.com/aoemedia-server/adapter/driving/restful/response"
	file2 "github.com/aoemedia-server/application/file"
	"github.com/aoemedia-server/config"
	"github.com/aoemedia-server/domain/file"
	"github.com/gin-gonic/gin"
	"time"
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
	metadata := file.NewMetadataBuilder().FileName(originalFileName).
		StorageDir(config.Inst().StorageFileRootDir()).Source(1).
		ModifiedTime(time.Now()).Build()
	domainFile, err := file.NewDomainFile(fileContent, metadata)

	storage := file2.NewFileStorage()

	_, saveErr := storage.SaveFile(domainFile)
	if saveErr != nil {
		response.SendInternalServerError(ctx, saveErr.Error())
		return
	}

	c.sendSuccessResponse(ctx, 0, originalFileName, fileContent.SizeInBytes, fileContent.HashValue)
}
