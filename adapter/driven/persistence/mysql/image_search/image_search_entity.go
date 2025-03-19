package image_search

import (
	"time"
)

// ImageSearch 图片搜索实体
type ImageSearch struct {
	// 主键ID
	ID int64 `gorm:"primaryKey"`
	// 用户ID
	UserId int64 `gorm:"column:user_id;not null"`
	// 文件ID
	FileId int64 `gorm:"column:file_id;not null"`
	// 来源 1:相机 2:微信
	Source uint8 `gorm:"not null"`
	// 修改时间
	ModifiedTime time.Time `gorm:"column:modified_time;not null"`
	// 文件完整路径
	FullPath string `gorm:"column:full_path;not null"`
	// 修改时间的年份
	Year int16 `gorm:"not null"`
	// 修改时间的月份
	Month uint8 `gorm:"not null"`
	// 修改时间的日期
	Day uint8 `gorm:"not null"`
	// 创建时间
	CreatedAt time.Time `gorm:"not null"`
}
