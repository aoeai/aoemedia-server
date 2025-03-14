package file

import (
	"gorm.io/gorm"
	"sync"

	mysqlfile "github.com/aoemedia-server/adapter/driven/persistence/mysql/file"
	"github.com/aoemedia-server/domain/file"
)

type Repository struct{}

var (
	instance *Repository
	once     sync.Once
)

func Inst() *Repository {
	once.Do(func() {
		instance = &Repository{}
	})
	return instance
}

func (r *Repository) Save(file *file.DomainFile, tx *gorm.DB) (fileId int64, error error) {
	return mysqlfile.Create(file, tx)
}
