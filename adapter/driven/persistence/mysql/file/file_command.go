package file

import (
	"github.com/aoemedia-server/domain/file"
	"gorm.io/gorm"
)

func Create(file *file.DomainFile, tx *gorm.DB) (fileId int64, error error) {
	entity := &File{}
	entity.Hash = file.HashValue
	entity.SizeInBytes = file.SizeInBytes
	entity.Filename = file.FileName
	entity.StorageDir = file.StorageDir
	entity.Source = file.Source
	entity.ModifiedTime = file.ModifiedTime

	tx = tx.Create(entity)

	return entity.ID, tx.Error
}
