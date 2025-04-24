package image

import (
	"fmt"
	"github.com/aoemedia-server/adapter/driven/repository/image_search"
	domainimagesearch "github.com/aoemedia-server/domain/image/image_search"
	"sync"
	"time"
)

var searcherInstance struct {
	searcher *Searcher
	once     sync.Once
}

type Searcher struct {
	imageSearchRepository domainimagesearch.Repository
}

func NewSearcher() *Searcher {
	searcherInstance.once.Do(func() {
		searcherInstance.searcher = newSearcher(image_search.Inst())
	})
	return searcherInstance.searcher
}

func newSearcher(imageSearchRepository domainimagesearch.Repository) *Searcher {
	searcherInstance.searcher = &Searcher{
		imageSearchRepository: imageSearchRepository,
	}
	return searcherInstance.searcher
}

func (s *Searcher) ImageList(params domainimagesearch.ImageSearchParams) ImageSearchResult {
	images := s.imageSearchRepository.ImageList(params)

	// 如果没有数据，直接返回空结果
	if len(images) == 0 {
		return ImageSearchResult{}
	}

	// 构建结果
	return ImageSearchResult{
		EarliestModifiedTime: findEarliestModifiedTime(images),
		GroupedDataList:      groupImagesByDate(images),
	}
}

// 寻找最早的修改时间
func findEarliestModifiedTime(images []domainimagesearch.ImageSearchResult) time.Time {
	if len(images) == 0 {
		return time.Time{}
	}

	earliestTime := images[0].ModifiedTime
	for _, img := range images {
		if img.ModifiedTime.Before(earliestTime) {
			earliestTime = img.ModifiedTime
		}
	}
	return earliestTime
}

// 按日期分组并保持输入顺序
func groupImagesByDate(images []domainimagesearch.ImageSearchResult) []GroupedData {
	// 按日期分组
	groupMap := make(map[string][]ImageInfo)
	// 记录分组的创建顺序，确保按输入顺序遍历
	groupOrder := make([]string, 0)

	for _, img := range images {
		title := formatDateTitle(img.Year, img.Month, img.Day)

		if _, exists := groupMap[title]; !exists {
			groupMap[title] = make([]ImageInfo, 0)
			groupOrder = append(groupOrder, title)
		}

		groupMap[title] = append(groupMap[title], ImageInfo{
			FullPath: img.FullPath,
		})
	}

	// 将分组数据转换为结果格式，保持输入顺序
	groupedData := make([]GroupedData, 0, len(groupMap))
	for _, title := range groupOrder {
		groupedData = append(groupedData, GroupedData{
			Title:  title,
			Images: groupMap[title],
		})
	}

	return groupedData
}

// 格式化日期标题
func formatDateTitle(year int16, month uint8, day uint8) string {
	return fmt.Sprintf("%d年%02d月%02d日", year, month, day)
}
