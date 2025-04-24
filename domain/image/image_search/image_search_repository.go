package image_search

import (
	domainimage "github.com/aoemedia-server/domain/image"
	"time"
)

type Repository interface {

	// SubscribeImageUploadedEvent 订阅发布图片已上传事件
	SubscribeImageUploadedEvent(event domainimage.ImageUploadedEvent)

	Save(imageSearch ImageSearch) (id int64, error error)

	// ImageList 获取图片列表
	ImageList(params ImageSearchParams) []ImageSearchResult
}

// ImageSearchParams 查询图片的参数
type ImageSearchParams struct {
	// 用户ID
	UserId int64
	// 修改时间
	ModifiedTime time.Time
	// 来源 1:相机 2:微信
	Source uint8
	// 每页记录数
	Limit int
}

type ImageSearchResult struct {
	// 文件ID
	FileId int64
	// 修改时间
	ModifiedTime time.Time
	// 文件完整路径
	FullPath string
	// 修改时间的年份
	Year int16
	// 修改时间的月份
	Month uint8
	// 修改时间的日期
	Day uint8
}
