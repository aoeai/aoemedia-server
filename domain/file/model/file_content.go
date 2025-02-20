package model

import (
	"crypto/sha256"
	"encoding/hex"
)

// FileContent 文件内容值对象
type FileContent struct {
	data        []byte
	hashValue   string
	sizeInBytes uint64
}

// NewFileContent 创建文件内容值对象
func NewFileContent(data []byte) *FileContent {
	return &FileContent{
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
func (fc *FileContent) SizeInBytes() uint64 {
	return fc.sizeInBytes
}

// Hash 文件哈希值
func (fc *FileContent) Hash() string {
	return fc.hashValue
}

// Data 获取文件内容数据
func (fc *FileContent) Data() []byte {
	return fc.data
}

// 获取文件的创建时间
