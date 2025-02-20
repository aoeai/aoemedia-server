package storage

import (
	"fmt"
	"github.com/aoemedia-server/application/storage/path"
	"github.com/aoemedia-server/config"
	"github.com/aoemedia-server/domain/image/model"
	imagestorage "github.com/aoemedia-server/domain/image/storage"
	"path/filepath"
	"time"
)

type ImageStorage struct {
	image *model.AoeImage
}

func NewImageStorage(image *model.AoeImage) (*ImageStorage, error) {
	if image == nil {
		return nil, fmt.Errorf("image 不能为空")
	}

	return &ImageStorage{image}, nil
}

// Save 存储图片
// 返回值:
//   - string: 文件存储的完整路径
//   - error: 存储过程中可能发生的错误
func (s *ImageStorage) Save(fileName string) (string, error) {
	fullDirPath := filepath.Join(config.Instance().FileStorage.ImageDir, createTimeOf(s.image))

	return s.save(fullDirPath, fileName)
}

func createTimeOf(image *model.AoeImage) string {
	if image.HasCreateTime() {
		return path.YearMonthOf(image.CreateTime())
	}

	return path.YearMonthOf(time.Now())
}

func (s *ImageStorage) save(fullDirPath, fileName string) (string, error) {
	// 创建图片存储器
	imageStorage, err := imagestorage.NewImageStorage(fullDirPath)
	if err != nil {
		return "", fmt.Errorf("创建图片存储器失败: %w", err)
	}

	// 保存文件并保持原始元数据
	return imageStorage.Save(s.image, fileName)
}
