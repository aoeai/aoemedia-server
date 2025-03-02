package file

import (
	"crypto/sha256"
	"encoding/hex"
)

// Content 文件内容值对象
type Content struct {
	data        []byte
	hashValue   string
	sizeInBytes uint64
}

// NewFileContent 创建文件内容值对象
func NewFileContent(data []byte) *Content {
	return &Content{
		data:        data,
		hashValue:   calculateHash(data),
		sizeInBytes: uint64(len(data)),
	}
}

// calculateHash 计算文件内容的SHA256哈希值
func calculateHash(data []byte) string {
	hash := sha256.New()
	hash.Write(data)
	return hex.EncodeToString(hash.Sum(nil))
}

// SizeInBytes 文件大小
func (c *Content) SizeInBytes() uint64 {
	return c.sizeInBytes
}

// Hash 文件哈希值
func (c *Content) Hash() string {
	return c.hashValue
}

// Data 获取文件内容数据
func (c *Content) Data() []byte {
	return c.data
}

// 获取文件的创建时间
