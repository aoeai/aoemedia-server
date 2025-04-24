package image_search

import (
	"testing"
	"time"

	domainimage "github.com/aoemedia-server/domain/image"
	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	tests := []struct {
		name    string
		event   domainimage.ImageUploadedEvent
		want    ImageSearch
		wantErr bool
	}{
		{
			name: "无效来源测试",
			event: domainimage.ImageUploadedEvent{
				UserId:         1,
				FileId:         1,
				Source:         0,
				ModifiedTime:   time.Now(),
				FullPathToFile: "/test.jpg",
			},
			want:    ImageSearch{},
			wantErr: true,
		},
		{
			name: "空路径测试",
			event: domainimage.ImageUploadedEvent{
				UserId:         1,
				FileId:         1,
				Source:         1,
				ModifiedTime:   time.Now(),
				FullPathToFile: "",
			},
			want:    ImageSearch{},
			wantErr: true,
		},
		{
			name: "零值时间测试",
			event: domainimage.ImageUploadedEvent{
				UserId:         1,
				FileId:         1,
				Source:         1,
				ModifiedTime:   time.Time{},
				FullPathToFile: "/test.jpg",
			},
			want:    ImageSearch{},
			wantErr: true,
		},
		{
			name: "正常转换测试",
			event: domainimage.ImageUploadedEvent{
				UserId:         123,
				FileId:         456,
				Source:         1,
				ModifiedTime:   time.Date(2024, 1, 15, 10, 30, 0, 0, time.Local),
				FullPathToFile: "/path/to/image.jpg",
			},
			wantErr: false,
			want: ImageSearch{
				UserId:       123,
				FileId:       456,
				Source:       1,
				ModifiedTime: time.Date(2024, 1, 15, 10, 30, 0, 0, time.Local),
				FullPath:     "/path/to/image.jpg",
				Year:         2024,
				Month:        1,
				Day:          15,
			},
		},
		{
			name: "跨年测试",
			event: domainimage.ImageUploadedEvent{
				UserId:         789,
				FileId:         101,
				Source:         2,
				ModifiedTime:   time.Date(2023, 12, 31, 23, 59, 59, 0, time.Local),
				FullPathToFile: "/path/to/another.jpg",
			},
			wantErr: false,
			want: ImageSearch{
				UserId:       789,
				FileId:       101,
				Source:       2,
				ModifiedTime: time.Date(2023, 12, 31, 23, 59, 59, 0, time.Local),
				FullPath:     "/path/to/another.jpg",
				Year:         2023,
				Month:        12,
				Day:          31,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := New(tt.event)
			if tt.wantErr {
				assert.Error(t, err)
				assert.Equal(t, tt.want, got)
				return
			}

			assert.NoError(t, err)
			assert.Equal(t, tt.want.UserId, got.UserId)
			assert.Equal(t, tt.want.FileId, got.FileId)
			assert.Equal(t, tt.want.Source, got.Source)
			assert.Equal(t, tt.want.ModifiedTime, got.ModifiedTime)
			assert.Equal(t, tt.want.FullPath, got.FullPath)
			assert.Equal(t, tt.want.Year, got.Year)
			assert.Equal(t, tt.want.Month, got.Month)
			assert.Equal(t, tt.want.Day, got.Day)
		})
	}
}
