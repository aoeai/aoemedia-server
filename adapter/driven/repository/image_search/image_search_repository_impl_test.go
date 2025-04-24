package image_search

import (
	"fmt"
	"github.com/aoemedia-server/common/converter"
	"testing"
	"time"

	domainimagesearch "github.com/aoemedia-server/domain/image/image_search"

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

func Test_ImageList(t *testing.T) {
	defaultParams := domainimagesearch.ImageSearchParams{
		UserId: 2,
		Limit:  10,
	}

	// 定义测试用例
	testCases := []struct {
		name           string
		params         domainimagesearch.ImageSearchParams
		mockReturn     []mysqlimagesearch.ImageSearch
		expectedLength int
		expectedChecks func(t *testing.T, result []domainimagesearch.ImageSearchResult)
	}{
		{
			name:           "空结果场景",
			params:         defaultParams,
			mockReturn:     []mysqlimagesearch.ImageSearch{},
			expectedLength: 0,
			expectedChecks: func(t *testing.T, result []domainimagesearch.ImageSearchResult) {
				assert.NotNil(t, result)
				assert.Empty(t, result)
			},
		},
		{
			name:           "1条记录场景",
			params:         defaultParams,
			mockReturn:     createMultipleImageSearchRecords(1),
			expectedLength: 1,
			expectedChecks: checkMultipleImageSearchRecords,
		},
		{
			name:           "3条记录场景",
			params:         defaultParams,
			mockReturn:     createMultipleImageSearchRecords(3),
			expectedLength: 3,
			expectedChecks: checkMultipleImageSearchRecords,
		}, {
			name:           "30条记录场景",
			params:         defaultParams,
			mockReturn:     createMultipleImageSearchRecords(30),
			expectedLength: 30,
			expectedChecks: checkMultipleImageSearchRecords,
		},
	}

	// 执行测试用例
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// 创建mock对象
			mockQuery := new(mysqlimagesearch.MockImageSearchQuery)

			// 设置模拟返回值
			mockQuery.On("ImageList", tc.params).Return(tc.mockReturn)

			// 创建仓库实例，注入mock
			repo := &Repository{
				imageSearchQuery: mockQuery,
			}

			// 执行测试
			result := repo.ImageList(tc.params)

			// 验证结果
			assert.Equal(t, tc.expectedLength, len(result))
			if tc.expectedChecks != nil {
				tc.expectedChecks(t, result)
			}

			// 验证mock方法被调用
			mockQuery.AssertCalled(t, "ImageList", tc.params)
		})
	}
}

// createMultipleImageSearchRecords 创建多条测试记录
func createMultipleImageSearchRecords(count int) []mysqlimagesearch.ImageSearch {
	records := make([]mysqlimagesearch.ImageSearch, count)
	for i := 0; i < count; i++ {
		records[i] = mysqlimagesearch.ImageSearch{
			FileId:       int64(200 + i),
			ModifiedTime: converter.StrToTime(fmt.Sprintf("2025-04-20 17:00:36.31170%d", i+1)),
			FullPath:     fmt.Sprintf("/path/to/image-%d.jpg", i+1),
		}
	}
	return records
}

// checkMultipleImageSearchRecords 检查多条记录的逻辑
func checkMultipleImageSearchRecords(t *testing.T, result []domainimagesearch.ImageSearchResult) {
	// 验证FileId是否递增
	for i := 0; i < len(result); i++ {
		assert.Equal(t, int64(200+i), result[i].FileId, "记录%d的FileId应为%d", i+1, 200+i)
	}

	// 验证ModifiedTime和FullPath的规律
	for i := 0; i < len(result); i++ {
		expectedTime := converter.StrToTime(fmt.Sprintf("2025-04-20 17:00:36.31170%d", i+1))
		expectedPath := fmt.Sprintf("/path/to/image-%d.jpg", i+1)

		assert.Equal(t, expectedTime, result[i].ModifiedTime, "记录%d的ModifiedTime不匹配", i+1)
		assert.Equal(t, expectedPath, result[i].FullPath, "记录%d的FullPath不匹配", i+1)
	}
}
