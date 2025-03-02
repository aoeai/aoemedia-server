package file

import "time"

type Metadata struct {
	fileName string
	// 存储路径
	storagePath string
	// 来源
	source int
	// 修改时间
	modifiedTime time.Time
}

func (m *Metadata) FileName() string {
	return m.fileName
}

func (m *Metadata) StoragePath() string {
	return m.storagePath
}

func (m *Metadata) Source() int {
	return m.source
}

func (m *Metadata) ModifiedTime() time.Time {
	return m.modifiedTime
}
