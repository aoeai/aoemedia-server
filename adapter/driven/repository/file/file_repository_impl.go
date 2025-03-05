package file

import (
	"github.com/aoemedia-server/adapter/driven/persistence/mysql/db"
	file2 "github.com/aoemedia-server/adapter/driven/persistence/mysql/file"
	"github.com/aoemedia-server/domain/file"
)

type Repository struct{}

func NewRepository() *Repository {
	return &Repository{}
}

func (r *Repository) Save(file *file.DomainFile) (int64, error) {
	return save(file)
}

func save(file *file.DomainFile) (int64, error) {
	entity := file2.File{}
	entity.Hash = file.HashValue
	entity.SizeInBytes = file.SizeInBytes
	entity.Filename = file.FileName
	entity.StorageDir = file.StorageDir
	entity.Source = file.Source
	entity.ModifiedTime = file.ModifiedTime

	tx := db.Inst().Create(&entity)

	return entity.ID, tx.Error
}
