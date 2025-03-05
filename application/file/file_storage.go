package file

import (
	localFileStorage "github.com/aoemedia-server/adapter/driven/persistence/local_storage/file"
	"github.com/aoemedia-server/domain/file"
	"sync"
)

type Storage struct {
	localStorage file.Storage
}

var (
	instance *Storage
	once     sync.Once
)

func NewFileStorage() *Storage {
	once.Do(func() {
		instance = &Storage{
			localStorage: localFileStorage.NewLocalFileStorage(),
		}
	})
	return instance
}

func (s *Storage) SaveFile(domainFile *file.DomainFile) (fullStoragePath string, err error) {
	return s.localStorage.Save(domainFile)
}
