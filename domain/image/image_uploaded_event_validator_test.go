package image

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestImageUploadedEvent_Validate(t *testing.T) {
	fixedTime := time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)
	tests := []struct {
		name    string
		subject *ImageUploadedEvent
		wantErr string
	}{
		{
			name: "所有字段都有效时，验证通过",
			subject: &ImageUploadedEvent{
				FileId:         1,
				UserId:         1,
				Source:         1,
				ModifiedTime:   fixedTime,
				FullPathToFile: "/path/to/file.jpg",
			},
			wantErr: "",
		},
		{
			name:    "event为nil时，返回错误：事件对象不能为空",
			subject: nil,
			wantErr: "事件对象不能为空",
		},
		{
			name: "FileId无效时，返回错误：文件ID不能为空或无效",
			subject: &ImageUploadedEvent{
				FileId:         0,
				UserId:         1,
				Source:         1,
				ModifiedTime:   fixedTime,
				FullPathToFile: "/path/to/file.jpg",
			},
			wantErr: "文件ID不能为空或无效",
		},
		{
			name: "UserId无效时，返回错误：用户ID不能为空或无效",
			subject: &ImageUploadedEvent{
				FileId:         1,
				UserId:         0,
				Source:         1,
				ModifiedTime:   fixedTime,
				FullPathToFile: "/path/to/file.jpg",
			},
			wantErr: "用户ID不能为空或无效",
		},
		{
			name: "Source为0时，返回错误：来源不能为空",
			subject: &ImageUploadedEvent{
				FileId:         1,
				UserId:         1,
				Source:         0,
				ModifiedTime:   fixedTime,
				FullPathToFile: "/path/to/file.jpg",
			},
			wantErr: "来源不能为空",
		},
		{
			name: "Source值无效时，返回错误：来源无效",
			subject: &ImageUploadedEvent{
				FileId:         1,
				UserId:         1,
				Source:         30,
				ModifiedTime:   fixedTime,
				FullPathToFile: "/path/to/file.jpg",
			},
			wantErr: "来源无效",
		},
		{
			name: "ModifiedTime为零值时，返回错误：修改时间不能为空",
			subject: &ImageUploadedEvent{
				FileId:         1,
				UserId:         1,
				Source:         1,
				ModifiedTime:   time.Time{},
				FullPathToFile: "/path/to/file.jpg",
			},
			wantErr: "修改时间不能为空",
		},
		{
			name: "FullPathToFile为空时，返回错误：文件路径不能为空",
			subject: &ImageUploadedEvent{
				FileId:         1,
				UserId:         1,
				Source:         1,
				ModifiedTime:   fixedTime,
				FullPathToFile: "",
			},
			wantErr: "文件路径不能为空",
		},
		{
			name: "FullPathToFile不是绝对路径时，返回错误：文件路径必须是绝对路径",
			subject: &ImageUploadedEvent{
				FileId:         1,
				UserId:         1,
				Source:         1,
				ModifiedTime:   fixedTime,
				FullPathToFile: "path/to/file.jpg",
			},
			wantErr: "文件路径必须是绝对路径",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.subject.Validate()
			if tt.wantErr == "" {
				assert.NoError(t, err)
			} else {
				assert.EqualError(t, err, tt.wantErr)
			}
		})
	}
}
