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
				data:        []byte("test data"),
				hashValue:   "1234567890123456789012345678901234567890123456789012345678901234",
				sizeInBytes: 9,
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
				data:        []byte{},
				hashValue:   "1234567890123456789012345678901234567890123456789012345678901234",
				sizeInBytes: 0,
			},
			wantErr: "文件内容大小必须大于0",
		},
		{
			name: "哈希值为空时，返回错误：文件内容哈希值不能为空",
			subject: &Content{
				data:        []byte("test data"),
				hashValue:   "",
				sizeInBytes: 9,
			},
			wantErr: "文件内容哈希值不能为空",
		},
		{
			name: "哈希值长度不是64时，返回错误：文件内容哈希值长度必须是64",
			subject: &Content{
				data:        []byte("test data"),
				hashValue:   "123456",
				sizeInBytes: 9,
			},
			wantErr: "文件内容哈希值长度必须是64",
		},
	}

	runValidationTests(t, tests)
}

func TestMetadata_validate(t *testing.T) {
	tests := []struct {
		name    string
		subject *Metadata
		wantErr string
	}{
		{
			name: "所有字段都有效时，验证通过",
			subject: &Metadata{
				fileName:     "test.jpg",
				storagePath:  "/path/to/storage",
				source:       1,
				modifiedTime: time.Now(),
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
				fileName:     "",
				storagePath:  "/path/to/storage",
				source:       1,
				modifiedTime: time.Now(),
			},
			wantErr: "文件名不能为空",
		},
		{
			name: "存储路径为空时，返回错误：存储路径不能为空",
			subject: &Metadata{
				fileName:     "test.jpg",
				storagePath:  "",
				source:       1,
				modifiedTime: time.Now(),
			},
			wantErr: "存储路径不能为空",
		},
		{
			name: "来源为0时，返回错误：文件来源不能为空",
			subject: &Metadata{
				fileName:     "test.jpg",
				storagePath:  "/path/to/storage",
				source:       0,
				modifiedTime: time.Now(),
			},
			wantErr: "文件来源不能为空",
		},
		{
			name: "来源无效时，返回错误：文件来源无效",
			subject: &Metadata{
				fileName:     "test.jpg",
				storagePath:  "/path/to/storage",
				source:       3,
				modifiedTime: time.Now(),
			},
			wantErr: "文件来源无效",
		},
		{
			name: "修改时间为零值时，返回错误：文件修改时间不能为空",
			subject: &Metadata{
				fileName:     "test.jpg",
				storagePath:  "/path/to/storage",
				source:       1,
				modifiedTime: time.Time{},
			},
			wantErr: "文件修改时间不能为空",
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
				content: &Content{
					data:        []byte("test data"),
					hashValue:   "1234567890123456789012345678901234567890123456789012345678901234",
					sizeInBytes: 9,
				},
				metadata: &Metadata{
					fileName:     "test.jpg",
					storagePath:  "/path/to/storage",
					source:       1,
					modifiedTime: time.Now(),
				},
			},
			wantErr: "",
		},
		{
			name: "文件内容无效时，返回错误",
			subject: &DomainFile{
				content: &Content{
					data:        []byte{},
					hashValue:   "1234567890123456789012345678901234567890123456789012345678901234",
					sizeInBytes: 0,
				},
				metadata: &Metadata{
					fileName:     "test.jpg",
					storagePath:  "/path/to/storage",
					source:       1,
					modifiedTime: time.Now(),
				},
			},
			wantErr: "文件内容大小必须大于0",
		},
		{
			name: "文件元数据无效时，返回错误",
			subject: &DomainFile{
				content: &Content{
					data:        []byte("test data"),
					hashValue:   "1234567890123456789012345678901234567890123456789012345678901234",
					sizeInBytes: 9,
				},
				metadata: &Metadata{
					fileName:     "",
					storagePath:  "/path/to/storage",
					source:       1,
					modifiedTime: time.Now(),
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
