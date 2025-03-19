package image

import (
	"path"

	"github.com/aoemedia-server/domain/image"
)

func NewImageUploadedEventPublishParams(domainImage *image.DomainImage, fileId int64, userId int64) *image.ImageUploadedEventPublishParams {
	return &image.ImageUploadedEventPublishParams{
		FileId:         fileId,
		UserId:         userId,
		Source:         domainImage.Source,
		ModifiedTime:   domainImage.ModifiedTime,
		FullPathToFile: path.Join(domainImage.StorageDir, domainImage.FileName),
	}
}

func newImageUploadedEvent(params image.ImageUploadedEventPublishParams) (image.ImageUploadedEvent, error) {
	event := image.ImageUploadedEvent{
		FileId:         params.FileId,
		UserId:         params.UserId,
		Source:         params.Source,
		ModifiedTime:   params.ModifiedTime,
		FullPathToFile: params.FullPathToFile,
	}

	err := event.Validate()
	if err != nil {
		return image.ImageUploadedEvent{}, err
	}

	return event, nil
}
