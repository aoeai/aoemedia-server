package image

import (
	repoimage "github.com/aoemedia-server/adapter/driven/repository/image"
	"github.com/aoemedia-server/domain/image"
	domainimage "github.com/aoemedia-server/domain/image"
	"sync"
)

type App struct {
	repository domainimage.Repository
}

var (
	instance *App
	once     sync.Once
)

func Inst() *App {
	once.Do(func() {
		instance = &App{repository: repoimage.Inst()}
	})
	return instance
}

// Upload 上传图片
//
// Params:
// - image: 图片
// - userId: 用户 id
//
// Returns:
// - error: 上传过程中可能发生的错误
func (app *App) Upload(image *image.DomainImage, userId int64) error {
	_, err := app.repository.Upload(image, userId)
	if err != nil {
		return err
	}

	return nil
}
