package image_search

import (
	"github.com/aoemedia-server/domain/image/image_search"
	"testing"
	"time"

	"github.com/aoemedia-server/adapter/driven/persistence/mysql/db"
	"github.com/stretchr/testify/assert"
)

func TestCreate(t *testing.T) {
	// 准备测试数据
	testTime := time.Now()
	tests := []struct {
		name        string
		imageSearch image_search.ImageSearch
		wantErr     bool
	}{
		{
			name: "正常场景-成功创建图片搜索记录",
			imageSearch: image_search.ImageSearch{
				UserId:       1,
				FileId:       160,
				Source:       1,
				ModifiedTime: testTime,
				FullPath:     "/path/to/image.jpg",
				Year:         int16(testTime.Year()),
				Month:        uint8(testTime.Month()),
				Day:          uint8(testTime.Day()),
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// 执行测试
			id, err := Create(tt.imageSearch)

			// 验证结果
			if tt.wantErr {
				assert.Error(t, err)
				return
			}

			// 验证成功场景
			assert.NoError(t, err)
			assert.Greater(t, id, int64(0))

			// 验证数据库中的记录
			var savedRecord ImageSearch
			result := db.Inst().First(&savedRecord, id)
			assert.NoError(t, result.Error)

			// 验证字段值
			assert.Equal(t, tt.imageSearch.UserId, savedRecord.UserId)
			assert.Equal(t, tt.imageSearch.FileId, savedRecord.FileId)
			assert.Equal(t, tt.imageSearch.Source, savedRecord.Source)
			assert.Equal(t, tt.imageSearch.ModifiedTime.Truncate(time.Millisecond), savedRecord.ModifiedTime.Truncate(time.Millisecond))
			assert.Equal(t, tt.imageSearch.FullPath, savedRecord.FullPath)
			assert.Equal(t, tt.imageSearch.Year, savedRecord.Year)
			assert.Equal(t, tt.imageSearch.Month, savedRecord.Month)
			assert.Equal(t, tt.imageSearch.Day, savedRecord.Day)
			assert.False(t, savedRecord.CreatedAt.IsZero())
		})
	}
}
