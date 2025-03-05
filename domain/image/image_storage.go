package image

import (
	"fmt"
	"github.com/aoemedia-server/domain/file"
	"github.com/sirupsen/logrus"
	"os"
	"path/filepath"
)

// Storage 图片文件存储器
type Storage struct {
	*file.LocalFileStorage
}

// NewImageStorage 创建图片文件存储器
func NewImageStorage(fullDirPath string) (*Storage, error) {
	localStorage, err := file.NewLocalFileStorage(fullDirPath)
	if err != nil {
		return nil, fmt.Errorf("创建本地存储器失败: %w", err)
	}

	return &Storage{LocalFileStorage: localStorage}, nil
}

// Save 存储图片文件并保持原始元数据
// 返回值:
//   - string: 文件存储的完整路径
//   - error: 存储过程中可能发生的错误
func (s *Storage) Save(domainImage *DomainImage, fileName string) (string, error) {
	// 保存文件
	relativePath, err := s.LocalFileStorage.Save(domainImage.FileContent(), fileName)
	if err != nil {
		return "", err
	}

	// 获取文件的完整路径
	fullPath := filepath.Join(s.LocalFileStorage.GetFullDirPath(), relativePath)

	// 设置文件的访问时间和修改时间为EXIF中的创建时间
	createTime := domainImage.CreateTime()
	if err := os.Chtimes(fullPath, createTime, createTime); err != nil {
		logrus.Warnf("%v 设置文件时间失败: %v", fileName, err)
	}
	logrus.Printf("%v 设置文件时间成功: %v", fileName, createTime)

	return fullPath, nil
}
