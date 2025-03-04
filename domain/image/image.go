package image

import (
	"fmt"
	"github.com/aoemedia-server/domain/file"
	"github.com/dsoprea/go-exif/v3"
	"time"
)

type DomainImage struct {
	fileContent   *file.Content
	createTime    time.Time
	hasCreateTime bool
}

func NewDomainImage(fc *file.Content) (*DomainImage, error) {
	if isImage := IsImage(fc); !isImage {
		return nil, fmt.Errorf("文件内容不是图片类型")
	}

	domainImage := &DomainImage{fileContent: fc, hasCreateTime: false}

	// 从EXIF中获取创建时间
	createTime, err := extractExifCreateTime(fc.Data())
	if err != nil {
		return domainImage, nil
	}
	if createTime.IsZero() {
		return domainImage, nil
	}
	domainImage.hasCreateTime = true
	domainImage.createTime = createTime

	return domainImage, nil
}

func (ai *DomainImage) FileContent() *file.Content {
	return ai.fileContent
}

func (ai *DomainImage) CreateTime() time.Time {
	return ai.createTime
}

func (ai *DomainImage) HasCreateTime() bool {
	return ai.hasCreateTime
}

// extractExifCreateTime 从图片数据中获取EXIF创建时间
func extractExifCreateTime(imageData []byte) (time.Time, error) {
	// 创建EXIF读取器
	rawExif, err := exif.SearchAndExtractExif(imageData)
	if err != nil {
		return time.Time{}, fmt.Errorf("读取EXIF数据失败: %w", err)
	}

	// 解析EXIF数据
	exifTags, _, err := exif.GetFlatExifData(rawExif, nil)
	if err != nil {
		return time.Time{}, fmt.Errorf("解析EXIF数据失败: %w", err)
	}

	// 尝试获取创建时间
	for _, ifd := range exifTags {
		if ifd.TagName == "DateTime" || ifd.TagName == "DateTimeOriginal" || ifd.TagName == "DateTimeDigitized" {
			strValue, ok := ifd.Value.(string)
			if !ok {
				continue
			}

			createTime, err := time.ParseInLocation("2006:01:02 15:04:05", strValue, time.Local)
			if err != nil {
				continue
			}

			return createTime, nil
		}
	}

	return time.Time{}, fmt.Errorf("未找到EXIF创建时间")
}
