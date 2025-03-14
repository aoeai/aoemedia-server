package image

import "github.com/aoemedia-server/adapter/driven/persistence/mysql/db"

func DeleteTestImageUploadRecordByFileId(fileId int64) {
	db.Inst().Delete(&ImageUploadRecord{}).Where("file_id = ?", fileId)
}
