package image_search

import (
	"github.com/aoemedia-server/common/eventbus"
	"github.com/aoemedia-server/domain/image/image_search"
	"sync"

	mysqlimagesearch "github.com/aoemedia-server/adapter/driven/persistence/mysql/image_search"
	domainimage "github.com/aoemedia-server/domain/image"
	domainimagesearch "github.com/aoemedia-server/domain/image/image_search"
	"github.com/sirupsen/logrus"
)

type Repository struct {
	imageSearchQuery mysqlimagesearch.ImageSearchQuery
}

var (
	instance *Repository
	once     sync.Once
)

func Inst() *Repository {
	once.Do(func() {
		instance = &Repository{
			imageSearchQuery: mysqlimagesearch.QueryInst(),
		}

		eventbus.Inst().Subscribe(domainimage.ImageUploadedEventType, func(data interface{}) {
			instance.SubscribeImageUploadedEvent(data.(domainimage.ImageUploadedEvent))
		})
	})
	return instance
}

func (r *Repository) SubscribeImageUploadedEvent(event domainimage.ImageUploadedEvent) {
	exist := r.imageSearchQuery.ExistByFileId(event.FileId)
	if exist {
		logrus.Warnf("订阅发布图片已上传事件-文件ID已存在: %v", event.FileId)
		return
	}

	imageSearch, err := image_search.New(event)
	if err != nil {
		logrus.Errorf("订阅发布图片已上传事件-创建 ImageSearch 失败: %v", err)
		return
	}

	_, err = r.Save(imageSearch)
	if err != nil {
		logrus.Errorf("订阅发布图片已上传事件-保存 ImageSearch 失败: %v", err)
	}
}

func (r *Repository) Save(imageSearch image_search.ImageSearch) (id int64, error error) {
	id, err := mysqlimagesearch.Create(imageSearch)
	if err != nil {
		return 0, err
	}

	logrus.Infof("保存 ImageSearch 成功 id:%v %v", id, imageSearch)
	return id, nil
}

func (r *Repository) ImageList(params domainimagesearch.ImageSearchParams) []domainimagesearch.ImageSearchResult {
	models := r.imageSearchQuery.ImageList(params)
	return convertToImageSearchResults(models)
}

// convertToImageSearchResults 将数据库模型转换为领域模型
func convertToImageSearchResults(models []mysqlimagesearch.ImageSearch) []domainimagesearch.ImageSearchResult {
	results := make([]domainimagesearch.ImageSearchResult, len(models))
	for i, model := range models {
		results[i] = convertToImageSearchResult(model)
	}
	return results
}

// convertToImageSearchResult 将单个数据库模型转换为领域模型
func convertToImageSearchResult(model mysqlimagesearch.ImageSearch) domainimagesearch.ImageSearchResult {
	return domainimagesearch.ImageSearchResult{
		FileId:       model.FileId,
		ModifiedTime: model.ModifiedTime,
		FullPath:     model.FullPath,
		Year:         model.Year,
		Month:        model.Month,
		Day:          model.Day,
	}
}
