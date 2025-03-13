package file

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestMetadataBuilder(t *testing.T) {
	testTime := time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC)

	tests := []struct {
		name     string
		builder  func() *MetadataBuilder
		expected *Metadata
	}{
		{
			name: "正常构建-所有字段都有效",
			builder: func() *MetadataBuilder {
				return NewMetadataBuilder().
					FileName("test.jpg").
					StorageDir("/storage/images").
					Source(1).
					ModifiedTime(testTime)
			},
			expected: &Metadata{
				FileName:     "test.jpg",
				StorageDir:   "/storage/images",
				Source:       1,
				ModifiedTime: testTime,
			},
		},
		{
			name: "构建失败-缺少必填字段",
			builder: func() *MetadataBuilder {
				return NewMetadataBuilder().
					FileName("").
					StorageDir("").
					Source(0)
			},
			expected: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.builder().Build()
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestMetadataBuilder_ChainCalls(t *testing.T) {
	testTime := time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC)
	builder := NewMetadataBuilder()

	// 测试链式调用返回的是同一个builder实例
	assert.Same(t, builder, builder.FileName("test.jpg"))
	assert.Same(t, builder, builder.StorageDir("/storage/images"))
	assert.Same(t, builder, builder.Source(1))
	assert.Same(t, builder, builder.ModifiedTime(testTime))

	// 验证最终构建的结果
	result := builder.Build()
	expected := &Metadata{
		FileName:     "test.jpg",
		StorageDir:   "/storage/images",
		Source:       1,
		ModifiedTime: testTime,
	}
	assert.Equal(t, expected, result)
}
