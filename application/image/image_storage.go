package image

import (
	"fmt"
	localFileStorage "github.com/aoemedia-server/adapter/driven/persistence/local_storage/file"
	"github.com/aoemedia-server/config"
	"github.com/aoemedia-server/domain/file"
	"github.com/aoemedia-server/domain/image"
	"path/filepath"
)

type Storage struct {
	localStorage file.Storage
	repository   file.Repository
}

func NewImageStorage(repository file.Repository) (*Storage, error) {
	if repository == nil {
		return nil, fmt.Errorf("repository 不能为空")
	}

	return &Storage{localStorage: localFileStorage.NewLocalFileStorage(), repository: repository}, nil
}

// Save 存储图片
// 返回值:
//   - int64: 文件ID
//   - string: 文件存储的完整目录
//   - error: 存储过程中可能发生的错误
func (s *Storage) Save(image *image.DomainImage) (int64, string, error) {
	fullDirPath := filepath.Join(config.Inst().Storage.ImageRootDir, createTimeOf(image))
	image.StorageDir = fullDirPath

	storageDir, err := s.localStorage.Save(image.DomainFile)
	if err != nil {
		return 0, "", err
	}

	id, err := s.repository.Save(image.DomainFile)

	return id, storageDir, err
}

func createTimeOf(image *image.DomainImage) string {
	return YearMonthOf(image.ModifiedTime)
}
