package image

import (
	"fmt"
	"github.com/aoemedia-server/domain/file"
	"time"
)

type DomainImage struct {
	*file.DomainFile
}

func NewDomainImage(domainFile *file.DomainFile) (*DomainImage, error) {
	if isImage := isImage(domainFile.Content.Data); !isImage {
		return nil, fmt.Errorf("文件内容不是图片类型")
	}

	domainImage := &DomainImage{DomainFile: domainFile}

	// 从EXIF中获取创建时间
	createTime, _ := extractExifCreateTime(domainFile.Data)
	if createTime.IsZero() {
		createTime = time.Now()
	}
	domainFile.ModifiedTime = createTime

	return domainImage, nil
}
