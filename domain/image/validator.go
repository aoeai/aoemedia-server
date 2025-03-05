package image

import (
	"github.com/aoemedia-server/domain/file"
	"github.com/h2non/filetype"
)

// isImage 判断文件内容是否为图片类型
func isImage(content *file.Content) bool {
	if content == nil {
		return false
	}

	data := content.Data
	if len(data) < 4 { // 至少需要4个字节才能判断
		return false
	}

	// 使用 filetype 库检测文件类型
	kind, err := filetype.Match(data)
	if err != nil {
		return false
	}

	// 检查是否为图片类型
	return kind.MIME.Type == "image"
}
