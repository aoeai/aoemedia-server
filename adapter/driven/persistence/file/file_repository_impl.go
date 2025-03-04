package file

import (
	"github.com/aoemedia-server/adapter/driven/persistence/db"
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
	entity := File{}
	entity.Hash = file.Content().Hash()
	entity.SizeInBytes = file.Content().SizeInBytes()
	entity.Filename = file.Metadata().FileName()
	entity.StoragePath = file.Metadata().StoragePath()
	entity.Source = file.Metadata().Source()
	entity.ModifiedTime = file.Metadata().ModifiedTime()

	tx := db.Inst().Create(&entity)

	return entity.ID, tx.Error
}
