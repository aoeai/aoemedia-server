package file

import (
	"crypto/sha256"
	"encoding/hex"
)

// Content 文件内容值对象
type Content struct {
	Data        []byte
	HashValue   string
	SizeInBytes int64
}

// NewFileContent 创建文件内容值对象
func NewFileContent(data []byte) *Content {
	return &Content{
		Data:        data,
		HashValue:   calculateHash(data),
		SizeInBytes: int64(len(data)),
	}
}

// calculateHash 计算文件内容的SHA256哈希值
func calculateHash(data []byte) string {
	hash := sha256.New()
	hash.Write(data)
	return hex.EncodeToString(hash.Sum(nil))
}
