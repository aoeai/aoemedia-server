package image_search

import (
	"fmt"
	"github.com/aoemedia-server/common/converter"
	domainimagesearch "github.com/aoemedia-server/domain/image/image_search"
	"testing"
	"time"

	"github.com/sirupsen/logrus"

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
	exists := QueryInst().ExistByFileId(imageSearch.FileId)
	assert.True(t, exists)

	// 清理测试数据
	result = db.Inst().Delete(imageSearch)
	assert.NoError(t, result.Error)
}

func shouldReturnFalseWhenFileNotExists(t *testing.T) {
	// 使用一个不存在的文件ID
	nonExistentFileId := int64(999)

	// 验证文件不存在
	exists := QueryInst().ExistByFileId(nonExistentFileId)
	assert.False(t, exists)
}

func initTestImages() {
	sql := `INSERT INTO image_search (user_id, file_id, source, modified_time, full_path, year, month, day, created_at)
    VALUES
    (201, 2100, 1, '2025-04-20 17:00:36.311701', '/path/to/image-1.jpg', 2025, 4, 20, '2025-04-20 17:00:36.312001'),
    (201, 2101, 1, '2025-04-20 17:00:36.311702', '/path/to/image-2.jpg', 2025, 4, 20, '2025-04-20 17:00:36.312002'),
    (201, 2102, 1, '2025-04-20 17:00:36.311703', '/path/to/image-3.jpg', 2025, 4, 20, '2025-04-20 17:00:36.312003'),
    (201, 2103, 1, '2025-04-20 17:00:36.311704', '/path/to/image-4.jpg', 2025, 4, 20, '2025-04-20 17:00:36.312004'),
    (201, 2104, 1, '2025-04-20 17:00:36.311705', '/path/to/image-5.jpg', 2025, 4, 20, '2025-04-20 17:00:36.312005'),
    (201, 2105, 1, '2025-04-20 17:00:36.311706', '/path/to/image-6.jpg', 2025, 4, 20, '2025-04-20 17:00:36.312006'),
    (201, 2106, 1, '2025-04-20 17:00:36.311707', '/path/to/image-7.jpg', 2025, 4, 20, '2025-04-20 17:00:36.312007'),
    (201, 2107, 1, '2025-04-20 17:00:36.311708', '/path/to/image-8.jpg', 2025, 4, 20, '2025-04-20 17:00:36.312008'),
    (201, 2108, 1, '2025-04-20 17:00:36.311709', '/path/to/image-9.jpg', 2025, 4, 20, '2025-04-20 17:00:36.312009'),
    (201, 2109, 1, '2025-04-20 17:00:36.311710', '/path/to/image-10.jpg', 2025, 4, 20, '2025-04-20 17:00:36.312010'),
    (201, 2200, 1, '2025-04-19 17:00:36.311701', '/path/to/image-2-1.jpg', 2025, 4, 19, '2025-04-20 17:00:36.312201'),
    (201, 2201, 1, '2025-04-19 17:00:36.311702', '/path/to/image-2-2.jpg', 2025, 4, 19, '2025-04-20 17:00:36.312202'),
    (201, 2300, 1, '2023-12-30 00:00:36.311701', '/path/to/image-3-1.jpg', 2023, 12, 30, '2025-04-20 17:00:36.312301'),
    (201, 2301, 1, '2023-12-30 00:00:36.311702', '/path/to/image-3-2.jpg', 2023, 12, 30, '2025-04-20 17:00:36.312302'),
    (201, 2302, 1, '2023-12-30 00:00:36.311703', '/path/to/image-3-3.jpg', 2023, 12, 30, '2025-04-20 17:00:36.312303')`

	result := db.Inst().Exec(sql)
	if result.Error != nil {
		logrus.Errorf("初始化测试图片数据失败: %v", result.Error)
	} else {
		logrus.Infof("成功初始化 %d 条测试图片数据", result.RowsAffected)
	}
}

func TestImageList(t *testing.T) {
	type imageListTestCase struct {
		name            string
		limit           int
		modifiedTime    time.Time
		expectedCount   int
		expectedFileIds []int64
	}

	// 定义测试用例
	testCases := []imageListTestCase{
		{
			name:            "当每页记录数=20，修改时间为未来时间，返回全部15条记录",
			limit:           20,
			modifiedTime:    parseTime("2025-04-20 18:00:36.312001"),
			expectedCount:   15,
			expectedFileIds: []int64{2109, 2108, 2107, 2106, 2105, 2104, 2103, 2102, 2101, 2100, 2201, 2200, 2302, 2301, 2300},
		},
		{
			name:            "当每页记录数=15，修改时间为未来时间，返回全部15条记录",
			limit:           15,
			modifiedTime:    parseTime("2025-04-20 18:00:36.312001"),
			expectedCount:   15,
			expectedFileIds: []int64{2109, 2108, 2107, 2106, 2105, 2104, 2103, 2102, 2101, 2100, 2201, 2200, 2302, 2301, 2300},
		},
		{
			name:            "当每页记录数=10，修改时间为未来时间，返回前10条记录",
			limit:           10,
			modifiedTime:    parseTime("2025-04-20 18:00:36.312001"),
			expectedCount:   10,
			expectedFileIds: []int64{2109, 2108, 2107, 2106, 2105, 2104, 2103, 2102, 2101, 2100},
		},
		{
			name:            "当每页记录数=10，修改时间等于第一批记录的最早时间，返回后5条记录",
			limit:           10,
			modifiedTime:    parseTime("2025-04-20 17:00:36.311701"),
			expectedCount:   5,
			expectedFileIds: []int64{2201, 2200, 2302, 2301, 2300},
		},
	}

	// 执行测试用例
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// 执行测试用例
			params := domainimagesearch.ImageSearchParams{
				UserId:       201,
				ModifiedTime: tc.modifiedTime,
				Source:       1,
				Limit:        tc.limit,
			}

			result := QueryInst().ImageList(params)

			// 使用assertThat风格验证结果
			assertImageList(t, result).
				hasCount(tc.expectedCount).
				hasFileIdsInOrder(tc.expectedFileIds)

			assert.Equal(t, tc.expectedCount, len(result), "返回记录数不匹配")
			for i, fileId := range tc.expectedFileIds {
				assert.Equal(t, fileId, result[i].FileId, fmt.Sprintf("第%d条记录的FileId应为%d", i+1, fileId))
			}
		})
	}
}

func parseTime(timeStr string) time.Time {
	return converter.StrToTime(timeStr)
}

type assertThat struct {
	t       *testing.T
	result  []ImageSearch
	message string
}

// hasCount 验证结果集的数量
func (a *assertThat) hasCount(expected int) *assertThat {
	assert.Equal(a.t, expected, len(a.result), a.message)
	return a
}

// hasFileIdsInOrder 验证结果集中的FileId顺序
func (a *assertThat) hasFileIdsInOrder(expectedFileIds []int64) *assertThat {
	for i, fileId := range expectedFileIds {
		assert.Equal(a.t, fileId, a.result[i].FileId, fmt.Sprintf("%s: 第%d条记录的FileId应为%d", a.message, i+1, fileId))
	}
	return a
}

// assertImageList 创建一个新的断言助手
func assertImageList(t *testing.T, result []ImageSearch) *assertThat {
	return &assertThat{t: t, result: result, message: "图片列表验证失败"}
}
