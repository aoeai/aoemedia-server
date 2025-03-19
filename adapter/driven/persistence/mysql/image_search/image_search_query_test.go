package image_search

import (
	"testing"
	"time"

	"github.com/aoemedia-server/adapter/driven/persistence/mysql/db"
	"github.com/stretchr/testify/assert"
)

func TestExistByFileId(t *testing.T) {
	t.Run("文件存在时返回true", shouldReturnTrueWhenFileExists)
	t.Run("文件不存在时返回false", shouldReturnFalseWhenFileNotExists)
}

func shouldReturnTrueWhenFileExists(t *testing.T) {
	// 准备测试数据
	testTime := time.Now()
	imageSearch := &ImageSearch{
		FileId:       200,
		ModifiedTime: testTime,
		FullPath:     "/path/to/image.jpg",
		Year:         int16(testTime.Year()),
		Month:        uint8(testTime.Month()),
		Day:          uint8(testTime.Day()),
	}

	// 保存测试数据
	result := db.Inst().Create(imageSearch)
	assert.NoError(t, result.Error)

	// 验证文件存在
	exists := ExistByFileId(imageSearch.FileId)
	assert.True(t, exists)

	// 清理测试数据
	result = db.Inst().Delete(imageSearch)
	assert.NoError(t, result.Error)
}

func shouldReturnFalseWhenFileNotExists(t *testing.T) {
	// 使用一个不存在的文件ID
	nonExistentFileId := int64(999)

	// 验证文件不存在
	exists := ExistByFileId(nonExistentFileId)
	assert.False(t, exists)
}
