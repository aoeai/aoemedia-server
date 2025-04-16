package image_search

import (
	domainimagesearch "github.com/aoemedia-server/domain/image/search_service"
	"testing"
	"time"

	"github.com/aoemedia-server/adapter/driven/persistence/mysql/db"

	mysqlimagesearch "github.com/aoemedia-server/adapter/driven/persistence/mysql/image_search"
	"github.com/stretchr/testify/assert"
)

func TestRepository_Save(t *testing.T) {
	t.Run("当保存成功时应该返回ID", shouldReturnIdWhenSaveSuccess)
	t.Run("保存成功后数据存储正确", shouldSaveDataCorrectly)
	t.Run("当文件ID已存在时应该返回错误", shouldReturnErrorWhenFileIdExist)
}

func shouldReturnIdWhenSaveSuccess(t *testing.T) {
	// 准备测试数据
	testTime := time.Now()
	imageSearch := domainimagesearch.ImageSearch{
		UserId:       1,
		FileId:       101,
		Source:       1,
		ModifiedTime: testTime,
		FullPath:     "/path/to/image.jpg",
		Year:         int16(testTime.Year()),
		Month:        uint8(testTime.Month()),
		Day:          uint8(testTime.Day()),
	}

	// 执行测试
	id, err := Inst().Save(imageSearch)

	// 验证结果
	assert.NoError(t, err)
	assert.Greater(t, id, int64(0))
}

func shouldSaveDataCorrectly(t *testing.T) {
	// 准备测试数据
	testTime := time.Now()
	imageSearch := domainimagesearch.ImageSearch{
		UserId:       1,
		FileId:       200,
		Source:       1,
		ModifiedTime: testTime,
		FullPath:     "/path/to/image.jpg",
		Year:         int16(testTime.Year()),
		Month:        uint8(testTime.Month()),
		Day:          uint8(testTime.Day()),
	}

	// 执行测试
	id, err := Inst().Save(imageSearch)
	assert.NoError(t, err)

	// 验证数据库中的记录
	var savedRecord mysqlimagesearch.ImageSearch
	result := db.Inst().First(&savedRecord, id)
	assert.NoError(t, result.Error)

	// 验证字段值
	assert.Equal(t, imageSearch.UserId, savedRecord.UserId)
	assert.Equal(t, imageSearch.FileId, savedRecord.FileId)
	assert.Equal(t, imageSearch.Source, savedRecord.Source)
	assert.Equal(t, imageSearch.ModifiedTime.Truncate(time.Millisecond), savedRecord.ModifiedTime.Truncate(time.Millisecond))
	assert.Equal(t, imageSearch.FullPath, savedRecord.FullPath)
	assert.Equal(t, imageSearch.Year, savedRecord.Year)
	assert.Equal(t, imageSearch.Month, savedRecord.Month)
	assert.Equal(t, imageSearch.Day, savedRecord.Day)
	assert.False(t, savedRecord.CreatedAt.IsZero())
}

func shouldReturnErrorWhenFileIdExist(t *testing.T) {
	// 准备测试数据
	testTime := time.Now()
	imageSearch := domainimagesearch.ImageSearch{
		UserId:       2,
		FileId:       100,
		Source:       1,
		ModifiedTime: testTime,
		FullPath:     "/path/to/image.jpg",
		Year:         int16(testTime.Year()),
		Month:        uint8(testTime.Month()),
		Day:          uint8(testTime.Day()),
	}

	// 先保存一条记录
	_, err := Inst().Save(imageSearch)
	assert.NoError(t, err)

	// 尝试保存相同file_id的记录
	imageSearch.FullPath = "/path/to/another/image.jpg"
	_, err = Inst().Save(imageSearch)

	// 验证是否返回错误
	assert.Error(t, err)
	// 验证是否包含预期的错误信息
	assert.Contains(t, err.Error(), "Duplicate entry")

	// 验证数据库中只有一条记录
	var count int64
	result := db.Inst().Model(&mysqlimagesearch.ImageSearch{}).Where("file_id = ?", imageSearch.FileId).Count(&count)
	assert.NoError(t, result.Error)
	assert.Equal(t, int64(1), count)
}
