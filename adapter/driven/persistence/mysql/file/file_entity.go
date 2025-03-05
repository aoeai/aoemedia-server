package file

import (
	"time"
)

// File 文件实体
type File struct {
	// 主键ID
	ID int64 `gorm:"primaryKey"`
	// 文件内容的哈希值
	Hash string `gorm:"not null"`
	// 文件大小 单位:字节
	SizeInBytes int64 `gorm:"not null"`
	// 文件名
	Filename string `gorm:"not null"`
	// 存储目录
	StorageDir string `gorm:"not null"`
	// 来源 1:相机 2:微信
	Source uint8 `gorm:"not null"`
	// 修改时间
	ModifiedTime time.Time `gorm:"not null"`
	// 创建时间
	CreatedAt time.Time `gorm:"not null"`
}
