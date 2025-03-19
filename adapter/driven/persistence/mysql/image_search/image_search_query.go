package image_search

import (
	"github.com/aoemedia-server/adapter/driven/persistence/mysql/db"
)

// ExistByFileId 根据文件ID判断是否存在
//
// Return:
//
// - bool: true 存在，false 不存在
func ExistByFileId(fileId int64) bool {
	var model ImageSearch
	db.Inst().Select("id").Where("file_id = ?", fileId).First(&model)

	return model.ID > 0
}
