package image_search

import (
	"github.com/aoemedia-server/common/eventbus"
	"sync"

	mysqlimagesearch "github.com/aoemedia-server/adapter/driven/persistence/mysql/image_search"
	domainimage "github.com/aoemedia-server/domain/image"
	domainimagesearch "github.com/aoemedia-server/domain/image_search"
	"github.com/sirupsen/logrus"
)

type Repository struct {
	repository *domainimagesearch.Repository
}

var (
	instance *Repository
	once     sync.Once
)

func Inst() *Repository {
	once.Do(func() {
		instance = &Repository{}

		eventbus.Inst().Subscribe(domainimage.ImageUploadedEventType, func(data interface{}) {
			instance.SubscribeImageUploadedEvent(data.(domainimage.ImageUploadedEvent))
		})
	})
	return instance
}

func (r *Repository) SubscribeImageUploadedEvent(event domainimage.ImageUploadedEvent) {
	exist := mysqlimagesearch.ExistByFileId(event.FileId)
	if exist {
		logrus.Warnf("订阅发布图片已上传事件-文件ID已存在: %v", event.FileId)
		return
	}

	imageSearch, err := domainimagesearch.New(event)
	if err != nil {
		logrus.Errorf("订阅发布图片已上传事件-创建 ImageSearch 失败: %v", err)
		return
	}

	_, err = r.Save(imageSearch)
	if err != nil {
		logrus.Errorf("订阅发布图片已上传事件-保存 ImageSearch 失败: %v", err)
	}
}

func (r *Repository) Save(imageSearch domainimagesearch.ImageSearch) (id int64, error error) {
	id, err := mysqlimagesearch.Create(imageSearch)
	if err != nil {
		return 0, err
	}

	logrus.Infof("保存 ImageSearch 成功 id:%v %v", id, imageSearch)
	return id, nil
}
