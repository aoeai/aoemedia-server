package image

type Repository interface {
	// Upload 上传图片
	//
	// Params:
	// - image: 图片
	// - userId: 用户 id
	//
	// Returns:
	// - UploadResult: 上传结果
	// - error: 上传过程中可能发生的错误
	Upload(image *DomainImage, userId int64) (*UploadResult, error)

	// PublishImageUploadedEvent 发布图片已上传事件
	//
	// Params:
	// - params: 图片上传事件发布参数
	//
	// Returns:
	// - ImageUploadedEvent: 图片已上传事件
	// - error: 发布过程中可能发生的错误
	PublishImageUploadedEvent(params *ImageUploadedEventPublishParams) error
}
