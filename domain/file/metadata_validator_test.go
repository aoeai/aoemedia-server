package file

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestMetadata_validate(t *testing.T) {
	tests := []struct {
		name    string
		subject *Metadata
		wantErr string
	}{
		{
			name: "所有字段都有效时，验证通过",
			subject: &Metadata{
				FileName:     "test.jpg",
				StorageDir:   "/path/to/storage",
				Source:       1,
				ModifiedTime: time.Now(),
			},
			wantErr: "",
		},
		{
			name:    "Metadata为nil时，返回错误：文件元数据不能为空",
			subject: nil,
			wantErr: "文件元数据不能为空",
		},
		{
			name: "文件名为空时，返回错误：文件名不能为空",
			subject: &Metadata{
				FileName:     "",
				StorageDir:   "/path/to/storage",
				Source:       1,
				ModifiedTime: time.Now(),
			},
			wantErr: "文件名不能为空",
		},
		{
			name: "存储路径为空时，返回错误：存储路径不能为空",
			subject: &Metadata{
				FileName:     "test.jpg",
				StorageDir:   "",
				Source:       1,
				ModifiedTime: time.Now(),
			},
			wantErr: "存储路径不能为空",
		},
		{
			name: "来源为0时，返回错误：来源不能为空",
			subject: &Metadata{
				FileName:     "test.jpg",
				StorageDir:   "/path/to/storage",
				Source:       0,
				ModifiedTime: time.Now(),
			},
			wantErr: "来源不能为空",
		},
		{
			name: "来源无效时，返回错误：来源无效",
			subject: &Metadata{
				FileName:     "test.jpg",
				StorageDir:   "/path/to/storage",
				Source:       3,
				ModifiedTime: time.Now(),
			},
			wantErr: "来源无效",
		},
		{
			name: "修改时间为零值时，返回错误：文件修改时间不能为空",
			subject: &Metadata{
				FileName:     "test.jpg",
				StorageDir:   "/path/to/storage",
				Source:       1,
				ModifiedTime: time.Time{},
			},
			wantErr: "文件修改时间不能为空",
		},
	}

	runValidationTests(t, tests)
}

func TestValidateSource(t *testing.T) {
	tests := []struct {
		name    string
		source  uint8
		wantErr string
	}{
		{
			name:    "来源为0时，返回错误：来源不能为空",
			source:  0,
			wantErr: "来源不能为空",
		},
		{
			name:    "来源为1(相机)时，验证通过",
			source:  1,
			wantErr: "",
		},
		{
			name:    "来源为2(微信)时，验证通过",
			source:  2,
			wantErr: "",
		},
		{
			name:    "来源为无效值时，返回错误：来源无效",
			source:  3,
			wantErr: "来源无效",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateSource(tt.source)
			if tt.wantErr == "" {
				assert.NoError(t, err)
			} else {
				assert.EqualError(t, err, tt.wantErr)
			}
		})
	}
}
