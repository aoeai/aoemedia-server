package file

import (
	"fmt"
	"github.com/aoemedia-server/domain/file"
	"github.com/sirupsen/logrus"
	"os"
	"path/filepath"
)

type LocalFileStorage struct {
}

func NewLocalFileStorage() *LocalFileStorage {
	return &LocalFileStorage{}
}

func (s *LocalFileStorage) Save(domainFile *file.DomainFile) (fullStoragePath string, err error) {
	return save(domainFile)
}

// Save 存储文件内容到本地文件系统
//
// fileContent: 要存储的文件内容对象，包含文件的二进制数据
// fileName: 存储的目标文件名
//
// Returns:
//
//   - string: 存储文件的完整目录
//   - error: 存储过程中可能发生的错误，包括：目录创建失败、文件已存在、写入失败等
func save(domainFile *file.DomainFile) (fullStorageDir string, err error) {
	storageFileRootDir := domainFile.StorageDir
	// 确保子目录存在
	if err := os.MkdirAll(storageFileRootDir, 0755); err != nil {
		return "", fmt.Errorf("创建文件子目录失败: %w", err)
	}

	// 构建完整的文件路径
	fullStorageDir = filepath.Join(storageFileRootDir, domainFile.FileName)

	// 如果文件已经存在，则返回错误
	if _, err := os.Stat(fullStorageDir); err == nil {
		return "", fmt.Errorf("文件已经存在: %s", fullStorageDir)
	}

	// 将文件内容写入到目标路径
	if err := os.WriteFile(fullStorageDir, domainFile.Data, 0644); err != nil {
		return "", fmt.Errorf("写入文件失败: %w", err)
	}

	// 返回相对于基础路径的存储路径
	relativePath, err := filepath.Rel(storageFileRootDir, fullStorageDir)
	if err != nil {
		return "", fmt.Errorf("获取相对路径失败: %w", err)
	}

	logrus.Infof("文件存储成功: fullFilePath:%s relativePath:%s", fullStorageDir, relativePath)

	return fullStorageDir, nil
}
