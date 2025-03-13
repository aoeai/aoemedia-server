package image

import (
	"path/filepath"
	"sync"

	localFileStorage "github.com/aoemedia-server/adapter/driven/persistence/local_storage/file"
	repofile "github.com/aoemedia-server/adapter/driven/repository/file"
	"github.com/aoemedia-server/common/eventbus"
	"github.com/aoemedia-server/config"
	"github.com/aoemedia-server/domain/file"
	"github.com/aoemedia-server/domain/image"
)

type Repository struct {
	fileLocalStorage file.Storage
	fileRepository   file.Repository
}

var (
	instance *Repository
	once     sync.Once
)

func Inst() *Repository {
	once.Do(func() {
		instance = &Repository{localFileStorage.NewLocalFileStorage(), repofile.NewRepository()}
	})
	return instance
}

func (r *Repository) Upload(domainImage *image.DomainImage, userId int64) (int64, error) {
	fileId, _, err := r.save(domainImage)
	if err != nil {
		return 0, err
	}

	return fileId, nil
}

// save 存储图片
// 返回值:
//   - int64: 文件ID
//   - string: 文件存储的完整目录
//   - error: 存储过程中可能发生的错误
func (r *Repository) save(image *image.DomainImage) (fileId int64, fullStorageDir string, err error) {
	fullDirPath := filepath.Join(config.Inst().Storage.ImageRootDir, createTimeOf(image))
	image.StorageDir = fullDirPath

	storageDir, err := r.fileLocalStorage.Save(image.DomainFile)
	if err != nil {
		return 0, "", err
	}

	id, err := r.fileRepository.Save(image.DomainFile)

	return id, storageDir, err
}

func createTimeOf(image *image.DomainImage) string {
	return YearMonthOf(image.ModifiedTime)
}

func (r *Repository) PublishImageUploadedEvent(params *image.ImageUploadedEventPublishParams) (image.ImageUploadedEvent, error) {
	event, err := newImageUploadedEvent(*params)
	if err != nil {
		return image.ImageUploadedEvent{}, err
	}

	eventbus.Inst().Publish(image.ImageUploadedEventType, event)

	return image.ImageUploadedEvent{}, nil
}
