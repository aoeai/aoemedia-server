package image

import (
	repoimage "github.com/aoemedia-server/adapter/driven/repository/image"
	repoimagesearch "github.com/aoemedia-server/adapter/driven/repository/image_search"
	"github.com/aoemedia-server/domain/image"
	domainimage "github.com/aoemedia-server/domain/image"
	domainimagesearch "github.com/aoemedia-server/domain/image/image_search"
	"github.com/sirupsen/logrus"
	"sync"
)

type App struct {
	imageRepository       domainimage.Repository
	imageSearchRepository domainimagesearch.Repository
}

var (
	uploaderInstance *App
	once             sync.Once
)

func NewUploader() *App {
	once.Do(func() {
		uploaderInstance = &App{imageRepository: repoimage.Inst(), imageSearchRepository: repoimagesearch.Inst()}
	})
	return uploaderInstance
}

// Upload 上传图片
//
// Params:
// - image: 图片
// - userId: 用户 id
//
// Returns:
// - error: 上传过程中可能发生的错误
func (app *App) Upload(image *image.DomainImage, userId int64) (result *domainimage.UploadResult, error error) {
	result, err := app.imageRepository.Upload(image, userId)
	if err != nil {
		return nil, err
	}

	params := repoimage.NewImageUploadedEventPublishParams(image, result.FileId, userId)
	app.publishEvent(params)

	return result, nil
}

func (app *App) publishEvent(params *image.ImageUploadedEventPublishParams) {
	err := app.imageRepository.PublishImageUploadedEvent(params)
	if err != nil {
		logrus.Errorf("发布图片已上传事件失败: %v", err)
	}
}
