package image

import "time"

// ImageUploadedEventPublishParams 图片上传事件发布参数
type ImageUploadedEventPublishParams struct {
	// UserId 用户ID
	UserId int64
	// FileId 文件ID
	FileId int64
	// 来源 1:相机 2:微信
	Source uint8
	// 修改时间
	ModifiedTime time.Time
	// 文件完整路径
	FullPathToFile string
}
