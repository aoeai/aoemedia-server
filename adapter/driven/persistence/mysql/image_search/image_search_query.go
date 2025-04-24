package image_search

import (
	"github.com/aoemedia-server/adapter/driven/persistence/mysql/db"
	domainimagesearch "github.com/aoemedia-server/domain/image/image_search"
	"sync"
)

type ImageSearchQuery interface {
	// ExistByFileId 根据文件ID判断是否存在
	//
	// Return:
	//
	// - bool: true 存在，false 不存在
	ExistByFileId(fileId int64) bool

	// ImageList 获取图片列表
	ImageList(params domainimagesearch.ImageSearchParams) []ImageSearch
}

type ImageSearchQueryInst struct{}

var (
	queryInstance ImageSearchQuery
	queryOnce     sync.Once
)

func QueryInst() ImageSearchQuery {
	queryOnce.Do(func() {
		queryInstance = new(ImageSearchQueryInst)
	})
	return queryInstance
}

// ExistByFileId 根据文件ID判断是否存在
//
// Return:
//
// - bool: true 存在，false 不存在
func (i *ImageSearchQueryInst) ExistByFileId(fileId int64) bool {
	var model ImageSearch
	db.Inst().Select("id").Where("file_id = ?", fileId).First(&model)

	return model.ID > 0
}

// ImageList 获取图片列表
func (i *ImageSearchQueryInst) ImageList(params domainimagesearch.ImageSearchParams) []ImageSearch {
	var models []ImageSearch
	db.Inst().Select("file_id,modified_time,full_path,year,month,day").
		Where("user_id = ?", params.UserId).
		Where("modified_time < ?", params.ModifiedTime).
		Where("source = ?", params.Source).
		Order("modified_time desc").
		Limit(params.Limit).
		Find(&models)

	return models
}
