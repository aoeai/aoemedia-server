package image_search

import (
	"github.com/aoemedia-server/application/image"
	"github.com/aoemedia-server/config"
	"strings"
)

type ListResult struct {
	// 最早的修改时间
	EarliestModifiedTime string
	// 按照年月日分组数据
	GroupedDataList []GroupedData
}

type GroupedData struct {
	// 标题：年 月 日
	Title  string
	Images []ImageInfo
}

type ImageInfo struct {
	// 图片访问地址
	Url string
}

func convertToListResult(svcResult image.ImageSearchResult) ListResult {
	return ListResult{
		EarliestModifiedTime: svcResult.EarliestModifiedTime.Format("2006-01-02 15:04:05.000000"),
		GroupedDataList:      convertToGroupedDataList(svcResult.GroupedDataList),
	}
}

func convertToGroupedDataList(groupedDataList []image.GroupedData) []GroupedData {
	if groupedDataList == nil {
		return []GroupedData{}
	}

	var result []GroupedData
	for _, data := range groupedDataList {
		result = append(result, GroupedData{
			Title:  data.Title,
			Images: convertToImageInfoList(data.Images),
		})
	}

	return result
}

func convertToImageInfoList(imageInfoList []image.ImageInfo) []ImageInfo {
	var result []ImageInfo
	for _, info := range imageInfoList {
		storage := config.Inst().Storage
		result = append(result, ImageInfo{
			Url: strings.Replace(info.FullPath, storage.ImageRootDir, storage.ImageURLPrefix, 1),
		})
	}
	return result
}
