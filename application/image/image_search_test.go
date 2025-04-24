package image

import (
	"fmt"
	"github.com/aoemedia-server/common/converter"
	"github.com/stretchr/testify/mock"
	"testing"
	"time"

	"github.com/aoemedia-server/domain/image/image_search"
	"github.com/stretchr/testify/assert"
)

func TestImageList(t *testing.T) {
	// 创建测试日期和时间
	dates := map[string]struct {
		year  int16
		month uint8
		day   uint8
		time  string
	}{
		"2025-05-21":  {2025, 5, 21, "2025-05-21 17:00:36.311701"},
		"2025-04-21":  {2025, 4, 21, "2025-04-21 17:00:36.311701"},
		"2025-04-21早": {2025, 4, 21, "2025-04-21 00:00:36.311701"},
		"2025-04-20晚": {2025, 4, 20, "2025-04-20 23:00:36.311701"},
		"2025-04-20":  {2025, 4, 20, "2025-04-20 17:00:36.311701"},
	}

	// 辅助函数：创建图片数据
	createImage := func(path string, dateKey string) image_search.ImageSearchResult {
		date := dates[dateKey]
		return image_search.ImageSearchResult{
			FullPath:     path,
			ModifiedTime: converter.StrToTime(date.time),
			Year:         date.year,
			Month:        date.month,
			Day:          date.day,
		}
	}

	// 辅助函数：创建图片结果
	createImageResult := func(path string) ImageInfo {
		return ImageInfo{FullPath: path}
	}

	// 辅助函数：创建分组
	createGroup := func(dateKey string, paths []string) GroupedData {
		date := dates[dateKey]
		title := fmt.Sprintf("%d年%02d月%02d日", date.year, date.month, date.day)

		images := make([]ImageInfo, 0, len(paths))
		for _, path := range paths {
			images = append(images, createImageResult(path))
		}

		return GroupedData{
			Title:  title,
			Images: images,
		}
	}

	// 定义测试用例
	testCases := []struct {
		name           string
		mockImages     []image_search.ImageSearchResult
		expectedResult ImageSearchResult
	}{
		{
			name:       "当数据库查询结果为空时，返回结果中的GroupedData为空",
			mockImages: []image_search.ImageSearchResult{},
			expectedResult: ImageSearchResult{
				EarliestModifiedTime: time.Time{},
				GroupedDataList:      []GroupedData{},
			},
		},
		{
			name: "当只有一条数据库记录时，返回结果正确",
			mockImages: []image_search.ImageSearchResult{
				createImage("/path/to/image1.jpg", "2025-04-20"),
			},
			expectedResult: ImageSearchResult{
				EarliestModifiedTime: converter.StrToTime(dates["2025-04-20"].time),
				GroupedDataList: []GroupedData{
					createGroup("2025-04-20", []string{"/path/to/image1.jpg"}),
				},
			},
		},
		{
			name: "当有2条日期不同的数据库记录时，返回结果正确",
			mockImages: []image_search.ImageSearchResult{
				createImage("/path/to/image2.jpg", "2025-04-21"),
				createImage("/path/to/image1.jpg", "2025-04-20"),
			},
			expectedResult: ImageSearchResult{
				EarliestModifiedTime: converter.StrToTime(dates["2025-04-20"].time),
				GroupedDataList: []GroupedData{
					createGroup("2025-04-21", []string{"/path/to/image2.jpg"}),
					createGroup("2025-04-20", []string{"/path/to/image1.jpg"}),
				},
			},
		},
		{
			name: "当前2条日期相同，后一条日期不同时，返回结果正确",
			mockImages: []image_search.ImageSearchResult{
				createImage("/path/to/image3.jpg", "2025-04-21"),
				createImage("/path/to/image2.jpg", "2025-04-21早"),
				createImage("/path/to/image1.jpg", "2025-04-20"),
			},
			expectedResult: ImageSearchResult{
				EarliestModifiedTime: converter.StrToTime(dates["2025-04-20"].time),
				GroupedDataList: []GroupedData{
					createGroup("2025-04-21", []string{"/path/to/image3.jpg", "/path/to/image2.jpg"}),
					createGroup("2025-04-20", []string{"/path/to/image1.jpg"}),
				},
			},
		},
		{
			name: "当第1条数据为一个日期，2、3条数据为一个日期时，返回结果正确",
			mockImages: []image_search.ImageSearchResult{
				createImage("/path/to/image3.jpg", "2025-05-21"),
				createImage("/path/to/image2.jpg", "2025-04-20晚"),
				createImage("/path/to/image1.jpg", "2025-04-20"),
			},
			expectedResult: ImageSearchResult{
				EarliestModifiedTime: converter.StrToTime(dates["2025-04-20"].time),
				GroupedDataList: []GroupedData{
					createGroup("2025-05-21", []string{"/path/to/image3.jpg"}),
					createGroup("2025-04-20", []string{"/path/to/image2.jpg", "/path/to/image1.jpg"}),
				},
			},
		},
	}

	// 执行测试用例
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// 创建mock对象
			mockRepository := new(image_search.MockImageSearchRepository)
			mockRepository.On("ImageList", mock.Anything).Return(tc.mockImages)

			// 创建searcher并执行测试
			searcher := newSearcher(mockRepository)
			result := searcher.ImageList(image_search.ImageSearchParams{})

			// 验证结果
			assert.Equal(t, tc.expectedResult.EarliestModifiedTime, result.EarliestModifiedTime, "最早修改时间不匹配")
			assert.Equal(t, len(tc.expectedResult.GroupedDataList), len(result.GroupedDataList), "分组数量不匹配")

			for i, group := range tc.expectedResult.GroupedDataList {
				assert.Equal(t, group.Title, result.GroupedDataList[i].Title, "分组标题不匹配，索引: %d", i)
				assert.Equal(t, len(group.Images), len(result.GroupedDataList[i].Images), "分组内图片��量不匹配，分组: %s", group.Title)
				for j, image := range group.Images {
					assert.Equal(t, image.FullPath, result.GroupedDataList[i].Images[j].FullPath,
						"图片路径不匹配，分组: %s, 图片索引: %d", group.Title, j)
				}
			}

			// 验证mock方法被调用
			mockRepository.AssertCalled(t, "ImageList", mock.Anything)
		})
	}
}
