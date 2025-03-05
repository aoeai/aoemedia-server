package image

import (
	"fmt"
	"github.com/aoemedia-server/domain/file"
	"github.com/dsoprea/go-exif/v3"
	"time"
)

type DomainImage struct {
	file        *file.DomainFile
	fileContent *file.Content
	createTime  time.Time
}

func NewDomainImage(fc *file.Content) (*DomainImage, error) {
	if isImage := isImage(fc); !isImage {
		return nil, fmt.Errorf("文件内容不是图片类型")
	}

	domainImage := &DomainImage{fileContent: fc}

	// 从EXIF中获取创建时间
	createTime, _ := extractExifCreateTime(fc.Data)
	if createTime.IsZero() {
		createTime = time.Now()
	}
	domainImage.createTime = createTime

	return domainImage, nil
}

func NewDomainImage1(fileName string, source uint8, fc *file.Content) (*DomainImage, error) {
	if isImage := isImage(fc); !isImage {
		return nil, fmt.Errorf("文件内容不是图片类型")
	}

	// 从EXIF中获取创建时间
	createTime, _ := extractExifCreateTime(fc.Data)
	if createTime.IsZero() {
		createTime = time.Now()
	}

	domainFile, err := newDomainFile(fileName, createTime, "", fc, source)
	if err != nil {
		return nil, err
	}

	domainImage := &DomainImage{file: domainFile}

	return domainImage, nil
}

func newDomainFile(fileName string, createTime time.Time, fullPath string, fileContent *file.Content, source uint8) (*file.DomainFile, error) {
	if createTime.IsZero() {
		createTime = time.Now()
	}
	metadata := file.NewMetadataBuilder().Source(source).FileName(fileName).StorageDir(fullPath).
		ModifiedTime(createTime).Build()
	imageFile, err := file.NewDomainFile(fileContent, metadata)
	return imageFile, err
}

func (di *DomainImage) FileContent() *file.Content {
	return di.fileContent
}

func (di *DomainImage) CreateTime() time.Time {
	return di.createTime
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
