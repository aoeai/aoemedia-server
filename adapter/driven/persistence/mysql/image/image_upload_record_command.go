package image

import (
	"gorm.io/gorm"
)

func Create(userId int64, fileId int64, tx *gorm.DB) (recordId int64, error error) {
	entity := &ImageUploadRecord{}
	entity.UserId = userId
	entity.FileId = fileId

	tx = tx.Create(entity)

	return entity.ID, tx.Error
}
