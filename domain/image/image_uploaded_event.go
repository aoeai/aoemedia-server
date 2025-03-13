package image

import (
	"time"
)

// ImageUploadedEvent 图片已上传事件
type ImageUploadedEvent struct {
	// 用户 id
	UserId int64
	// 文件 id
	FileId int64
	// 来源 1:相机 2:微信
	Source uint8
	// 修改时间
	ModifiedTime time.Time
	// 文件完整路径
	FullPathToFile string
}
