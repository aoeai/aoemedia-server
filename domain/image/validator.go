package image

import (
	"github.com/h2non/filetype"
)

// isImage 判断文件内容是否为图片类型
func isImage(data []byte) bool {
	if data == nil {
		return false
	}

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
