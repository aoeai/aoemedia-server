package image

import "time"

type ImageSearchResult struct {
	// 最早的修改时间
	EarliestModifiedTime time.Time
	// 按照年月日分组数据
	GroupedDataList []GroupedData
}

type GroupedData struct {
	// 标题：年 月 日
	Title  string
	Images []ImageInfo
}

type ImageInfo struct {
	// 文件完整路径
	FullPath string
}
