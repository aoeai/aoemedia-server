package storage

import (
	"fmt"
	"github.com/aoemedia-server/config"
	"github.com/aoemedia-server/domain/file/model"
	"github.com/aoemedia-server/domain/file/storage"
)

type FileStorage struct {
	fileContent *model.FileContent
}

func NewFileStorage(fileContent *model.FileContent) (*FileStorage, error) {
	if fileContent == nil {
		return nil, fmt.Errorf("FileContent 不能为空")
	}

	return &FileStorage{fileContent}, nil
}

func (s *FileStorage) Save(filename string) (string, error) {
	fullDirPath := config.Instance().RootDirPath()

	return s.save(fullDirPath, filename)
}

func (s *FileStorage) save(fullDirPath, filename string) (string, error) {
	localStorage, err := storage.NewLocalFileStorage(fullDirPath)
	if err != nil {
		return "", err
	}

	return localStorage.Save(s.fileContent, filename)
}
