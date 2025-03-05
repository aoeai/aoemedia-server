package image

import (
	"fmt"
	"path/filepath"
	"time"

	"github.com/aoemedia-server/config"
	"github.com/aoemedia-server/domain/file"
	"github.com/aoemedia-server/domain/image"
)

type Storage struct {
	image      *image.DomainImage
	repository file.Repository
}

func NewImageStorage(image *image.DomainImage, repository file.Repository) (*Storage, error) {
	if image == nil {
		return nil, fmt.Errorf("image 不能为空")
	}
	if repository == nil {
		return nil, fmt.Errorf("repository 不能为空")
	}

	return &Storage{image: image, repository: repository}, nil
}

// Save 存储图片
// 返回值:
//   - int64: 文件ID
//   - string: 文件存储的完整路径
//   - error: 存储过程中可能发生的错误
func (s *Storage) Save(fileName string) (int64, string, error) {
	fullDirPath := filepath.Join(config.Inst().Storage.ImageRootDir, createTimeOf(s.image))

	fullPath, err := s.save(fullDirPath, fileName)
	if err != nil {
		return 0, "", err
	}

	imageFile, err := newDomainFile(fileName, s, fullPath)
	if err != nil {
		return 0, "", err
	}

	id, err := s.repository.Save(imageFile)

	return id, fullPath, err
}

func createTimeOf(image *image.DomainImage) string {
	return YearMonthOf(image.CreateTime())
}

func (s *Storage) save(fullDirPath, fileName string) (string, error) {
	// 创建图片存储器
	imageStorage, err := image.NewImageStorage(fullDirPath)
	if err != nil {
		return "", fmt.Errorf("创建图片存储器失败: %w", err)
	}

	// 保存文件并保持原始元数据
	return imageStorage.Save(s.image, fileName)
}

func newDomainFile(fileName string, s *Storage, fullPath string) (*file.DomainFile, error) {
	createTime := s.image.CreateTime()
	if createTime.IsZero() {
		createTime = time.Now()
	}
	metadata := file.NewMetadataBuilder().Source(1).FileName(fileName).StorageDir(fullPath).
		ModifiedTime(createTime).Build()
	imageFile, err := file.NewDomainFile(s.image.FileContent(), metadata)
	return imageFile, err
}
