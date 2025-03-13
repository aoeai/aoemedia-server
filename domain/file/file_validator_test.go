package file

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestContent_validate(t *testing.T) {
	tests := []struct {
		name    string
		subject *Content
		wantErr string
	}{
		{
			name: "所有字段都有效时，验证通过",
			subject: &Content{
				Data:        []byte("test Data"),
				HashValue:   "1234567890123456789012345678901234567890123456789012345678901234",
				SizeInBytes: 9,
			},
			wantErr: "",
		},
		{
			name:    "Content为nil时，返回错误：文件内容不能为空",
			subject: nil,
			wantErr: "文件内容不能为空",
		},
		{
			name: "文件大小为0时，返回错误：文件内容大小必须大于0",
			subject: &Content{
				Data:        []byte{},
				HashValue:   "1234567890123456789012345678901234567890123456789012345678901234",
				SizeInBytes: 0,
			},
			wantErr: "文件内容大小必须大于0",
		},
		{
			name: "哈希值为空时，返回错误：文件内容哈希值不能为空",
			subject: &Content{
				Data:        []byte("test Data"),
				HashValue:   "",
				SizeInBytes: 9,
			},
			wantErr: "文件内容哈希值不能为空",
		},
		{
			name: "哈希值长度不是64时，返回错误：文件内容哈希值长度必须是64",
			subject: &Content{
				Data:        []byte("test Data"),
				HashValue:   "123456",
				SizeInBytes: 9,
			},
			wantErr: "文件内容哈希值长度必须是64",
		},
	}

	runValidationTests(t, tests)
}

func TestFile_validate(t *testing.T) {
	tests := []struct {
		name    string
		subject *DomainFile
		wantErr string
	}{
		{
			name: "所有字段都有效时，验证通过",
			subject: &DomainFile{
				Content: &Content{
					Data:        []byte("test Data"),
					HashValue:   "1234567890123456789012345678901234567890123456789012345678901234",
					SizeInBytes: 9,
				},
				Metadata: &Metadata{
					FileName:     "test.jpg",
					StorageDir:   "/path/to/storage",
					Source:       1,
					ModifiedTime: time.Now(),
				},
			},
			wantErr: "",
		},
		{
			name: "文件内容无效时，返回错误",
			subject: &DomainFile{
				Content: &Content{
					Data:        []byte{},
					HashValue:   "1234567890123456789012345678901234567890123456789012345678901234",
					SizeInBytes: 0,
				},
				Metadata: &Metadata{
					FileName:     "test.jpg",
					StorageDir:   "/path/to/storage",
					Source:       1,
					ModifiedTime: time.Now(),
				},
			},
			wantErr: "文件内容大小必须大于0",
		},
		{
			name: "文件元数据无效时，返回错误",
			subject: &DomainFile{
				Content: &Content{
					Data:        []byte("test Data"),
					HashValue:   "1234567890123456789012345678901234567890123456789012345678901234",
					SizeInBytes: 9,
				},
				Metadata: &Metadata{
					FileName:     "",
					StorageDir:   "/path/to/storage",
					Source:       1,
					ModifiedTime: time.Now(),
				},
			},
			wantErr: "文件名不能为空",
		},
	}

	runValidationTests(t, tests)
}

type validator interface {
	validate() error
}

func runValidationTests[T validator](t *testing.T, tests []struct {
	name    string
	subject T
	wantErr string
}) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.subject.validate()
			if tt.wantErr == "" {
				assert.NoError(t, err)
			} else {
				assert.EqualError(t, err, tt.wantErr)
			}
		})
	}
}
