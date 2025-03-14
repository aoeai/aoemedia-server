package file

import "gorm.io/gorm"

type Repository interface {
	Save(file *DomainFile, tx *gorm.DB) (fileId int64, error error)
}
