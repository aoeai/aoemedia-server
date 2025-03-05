package file

import "time"

type Metadata struct {
	FileName string
	// 存储目录
	StorageDir string
	// 来源
	Source uint8
	// 修改时间
	ModifiedTime time.Time
}
