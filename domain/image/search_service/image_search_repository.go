package search_service

import domainimage "github.com/aoemedia-server/domain/image"

type Repository interface {

	// SubscribeImageUploadedEvent 订阅发布图片已上传事件
	SubscribeImageUploadedEvent(event domainimage.ImageUploadedEvent)

	Save(imageSearch ImageSearch) (id int64, error error)
}
