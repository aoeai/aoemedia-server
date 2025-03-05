package file

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"os"
	"path/filepath"
)

// LocalFileStorage 本地文件存储器
type LocalFileStorage struct {
	fullDirPath string
}

// NewLocalFileStorage 创建本地文件存储器
func NewLocalFileStorage(fullDirPath string) (*LocalFileStorage, error) {
	// 确保存储基础路径存在
	if err := os.MkdirAll(fullDirPath, 0755); err != nil {
		return nil, fmt.Errorf("创建存储目录失败: %w", err)
	}

	return &LocalFileStorage{fullDirPath: fullDirPath}, nil
}

// Save 存储文件内容到本地文件系统
// fileContent: 要存储的文件内容对象，包含文件的二进制数据
// fileName: 存储的目标文件名
// 返回值:
//   - string: 文件存储后相对于存储根目录的相对路径
//   - error: 存储过程中可能发生的错误，包括：目录创建失败、文件已存在、写入失败等
func (s *LocalFileStorage) Save(fileContent *Content, fileName string) (string, error) {
	// 确保子目录存在
	if err := os.MkdirAll(s.fullDirPath, 0755); err != nil {
		return "", fmt.Errorf("创建文件子目录失败: %w", err)
	}

	// 构建完整的文件路径
	fullFilePath := filepath.Join(s.fullDirPath, fileName)

	// 如果文件已经存在，则返回错误
	if _, err := os.Stat(fullFilePath); err == nil {
		return "", fmt.Errorf("文件已经存在: %s", fullFilePath)
	}

	// 将文件内容写入到目标路径
	if err := os.WriteFile(fullFilePath, fileContent.Data, 0644); err != nil {
		return "", fmt.Errorf("写入文件失败: %w", err)
	}

	// 返回相对于基础路径的存储路径
	relativePath, err := filepath.Rel(s.fullDirPath, fullFilePath)
	if err != nil {
		return "", fmt.Errorf("获取相对路径失败: %w", err)
	}

	logrus.Infof("文件存储成功: fullFilePath:%s relativePath:%s", fullFilePath, relativePath)

	return relativePath, nil
}

// GetFullDirPath 获取存储目录的完整路径
func (s *LocalFileStorage) GetFullDirPath() string {
	return s.fullDirPath
}
