package image

import (
	"fmt"
	"path/filepath"

	"github.com/aoemedia-server/domain/file"
)

func (event *ImageUploadedEvent) Validate() error {
	if event == nil {
		return fmt.Errorf("事件对象不能为空")
	}

	// image.ImageUploadedEvent中的所有字段不能为空
	if event.FileId <= 0 {
		return fmt.Errorf("文件ID不能为空或无效")
	}
	if event.UserId <= 0 {
		return fmt.Errorf("用户ID不能为空或无效")
	}
	if err := file.ValidateSource(event.Source); err != nil {
		return err
	}
	if event.ModifiedTime.IsZero() {
		return fmt.Errorf("修改时间不能为空")
	}
	if event.FullPathToFile == "" {
		return fmt.Errorf("文件路径不能为空")
	}
	// 验证文件路径格式
	if !filepath.IsAbs(event.FullPathToFile) {
		return fmt.Errorf("文件路径必须是绝对路径")
	}

	return nil
}
