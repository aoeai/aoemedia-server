package image

import (
	"fmt"
	"github.com/aoemedia-server/domain/file"
	"github.com/sirupsen/logrus"
	"os"
	"path/filepath"
)

// ImageStorage 图片文件存储器
type ImageStorage struct {
	*file.LocalFileStorage
}

// NewImageStorage 创建图片文件存储器
func NewImageStorage(fullDirPath string) (*ImageStorage, error) {
	localStorage, err := file.NewLocalFileStorage(fullDirPath)
	if err != nil {
		return nil, fmt.Errorf("创建本地存储器失败: %w", err)
	}

	return &ImageStorage{LocalFileStorage: localStorage}, nil
}

// Save 存储图片文件并保持原始元数据
// 返回值:
//   - string: 文件存储的完整路径
//   - error: 存储过程中可能发生的错误
func (s *ImageStorage) Save(aoeImage *AoeImage, fileName string) (string, error) {
	// 保存文件
	relativePath, err := s.LocalFileStorage.Save(aoeImage.FileContent(), fileName)
	if err != nil {
		return "", err
	}

	// 获取文件的完整路径
	fullPath := filepath.Join(s.LocalFileStorage.GetFullDirPath(), relativePath)

	hasCreateTime := aoeImage.HasCreateTime()
	if !hasCreateTime {
		return fullPath, nil
	}

	// 设置文件的访问时间和修改时间为EXIF中的创建时间
	createTime := aoeImage.CreateTime()
	if err := os.Chtimes(fullPath, createTime, createTime); err != nil {
		logrus.Warnf("%v 设置文件时间失败: %v", fileName, err)
	}
	logrus.Printf("%v 设置文件时间成功: %v", fileName, createTime)

	return fullPath, nil
}
