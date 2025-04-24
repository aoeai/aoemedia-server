package image_search

import (
	"fmt"
	domainimage "github.com/aoemedia-server/domain/image"
	"strings"
)

func validateEvent(event domainimage.ImageUploadedEvent) error {
	if event == (domainimage.ImageUploadedEvent{}) {
		return fmt.Errorf("ImageUploadedEvent 不能为空")
	}

	if event.UserId <= 0 {
		return fmt.Errorf("UserId 必须大于0")
	}

	if event.FileId <= 0 {
		return fmt.Errorf("FileId 必须大于0")
	}

	if event.Source < 1 {
		return fmt.Errorf("source 必须大于0")
	}

	if event.ModifiedTime.IsZero() {
		return fmt.Errorf("ModifiedTime 不能为空")
	}

	if strings.TrimSpace(event.FullPathToFile) == "" {
		return fmt.Errorf("FullPathToFile 不能为空")
	}

	return nil
}
