package image_search

import (
	"testing"
	"time"

	domainimage "github.com/aoemedia-server/domain/image"
	"github.com/stretchr/testify/assert"
)

func TestValidateEvent(t *testing.T) {
	fixedTime := time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)
	tests := []struct {
		name    string
		event   domainimage.ImageUploadedEvent
		wantErr string
	}{
		{
			name: "所有字段都有效时，验证通过",
			event: domainimage.ImageUploadedEvent{
				FileId:         1,
				UserId:         1,
				Source:         1,
				ModifiedTime:   fixedTime,
				FullPathToFile: "/path/to/file.jpg",
			},
			wantErr: "",
		},
		{
			name:    "空事件时，返回错误",
			event:   domainimage.ImageUploadedEvent{},
			wantErr: "ImageUploadedEvent 不能为空",
		},
		{
			name: "UserId无效时，返回错误",
			event: domainimage.ImageUploadedEvent{
				FileId:         1,
				UserId:         0,
				Source:         1,
				ModifiedTime:   fixedTime,
				FullPathToFile: "/path/to/file.jpg",
			},
			wantErr: "UserId 必须大于0",
		},
		{
			name: "FileId无效时，返回错误",
			event: domainimage.ImageUploadedEvent{
				FileId:         0,
				UserId:         1,
				Source:         1,
				ModifiedTime:   fixedTime,
				FullPathToFile: "/path/to/file.jpg",
			},
			wantErr: "FileId 必须大于0",
		},
		{
			name: "Source无效时，返回错误",
			event: domainimage.ImageUploadedEvent{
				FileId:         1,
				UserId:         1,
				Source:         0,
				ModifiedTime:   fixedTime,
				FullPathToFile: "/path/to/file.jpg",
			},
			wantErr: "source 必须大于0",
		},
		{
			name: "ModifiedTime为零值时，返回错误",
			event: domainimage.ImageUploadedEvent{
				FileId:         1,
				UserId:         1,
				Source:         1,
				ModifiedTime:   time.Time{},
				FullPathToFile: "/path/to/file.jpg",
			},
			wantErr: "ModifiedTime 不能为空",
		},
		{
			name: "FullPathToFile为空时，返回错误",
			event: domainimage.ImageUploadedEvent{
				FileId:         1,
				UserId:         1,
				Source:         1,
				ModifiedTime:   fixedTime,
				FullPathToFile: "",
			},
			wantErr: "FullPathToFile 不能为空",
		},
		{
			name: "FullPathToFile只包含空格时，返回错误",
			event: domainimage.ImageUploadedEvent{
				FileId:         1,
				UserId:         1,
				Source:         1,
				ModifiedTime:   fixedTime,
				FullPathToFile: "   ",
			},
			wantErr: "FullPathToFile 不能为空",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := validateEvent(tt.event)
			if tt.wantErr == "" {
				assert.NoError(t, err)
			} else {
				assert.EqualError(t, err, tt.wantErr)
			}
		})
	}
}
