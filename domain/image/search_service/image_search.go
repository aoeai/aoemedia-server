package search_service

import (
	"time"

	domainimage "github.com/aoemedia-server/domain/image"
	"github.com/sirupsen/logrus"
)

type ImageSearch struct {
	// 用户ID
	UserId int64
	// 文件ID
	FileId int64
	// 来源 1:相机 2:微信
	Source uint8
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

func New(event domainimage.ImageUploadedEvent) (ImageSearch, error) {
	if err := validateEvent(event); err != nil {
		logrus.Errorf("验证 ImageUploadedEvent 失败: %v", err)
		return ImageSearch{}, err
	}

	return ImageSearch{
		UserId:       event.UserId,
		FileId:       event.FileId,
		Source:       event.Source,
		ModifiedTime: event.ModifiedTime,
		FullPath:     event.FullPathToFile,
		Year:         int16(event.ModifiedTime.Year()),
		Month:        uint8(event.ModifiedTime.Month()),
		Day:          uint8(event.ModifiedTime.Day()),
	}, nil
}
