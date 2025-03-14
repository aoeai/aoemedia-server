package image

import (
	"time"
)

// ImageUploadRecord 图片上传记录实体
type ImageUploadRecord struct {
	// 主键ID
	ID int64 `gorm:"primaryKey"`
	// 用户ID
	UserId int64 `gorm:"column:user_id;not null"`
	// 文件ID
	FileId int64 `gorm:"column:file_id;not null"`
	// 创建时间
	CreatedAt time.Time `gorm:"not null"`
}
