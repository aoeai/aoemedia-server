package image

import (
	"github.com/aoemedia-server/common/testconst"
	"github.com/aoemedia-server/domain/file"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func Test_extractExifCreateTime(t *testing.T) {
	t.Run("当文件中包含EXIF时提取出正确的修改时间", shouldExtractCorrectModTimeWhenFileIsWebp)
	t.Run("当文件中不包含EXIF时返回零时间", shouldReturnZeroTimeWhenFileHasNoExif)
}

func shouldExtractCorrectModTimeWhenFileIsWebp(t *testing.T) {
	type args struct {
		testFileName string
		expectedTime time.Time
	}

	tests := []args{
		{testFileName: testconst.Jpg, expectedTime: time.Date(2024, time.May, 15, 8, 59, 5, 0, time.Local)},
	}

	for _, arg := range tests {
		t.Run(arg.testFileName, func(t *testing.T) {
			testFilePath := file.DomainFileTestdataPath(arg.testFileName)

			// 读取测试文件内容
			fileContent := file.NewTestFileContent(t, testFilePath)

			// 提取EXIF创建时间
			actualTime, err := extractExifCreateTime(fileContent.Data)

			// 验证结果
			assert.NoError(t, err)
			assert.Equal(t, arg.expectedTime, actualTime, "EXIF创建时间应该是 %v，但实际是 %v",
				arg.expectedTime.Format("2006-01-02 15:04:05"), actualTime.Format("2006-01-02 15:04:05"))
		})
	}
}

func shouldReturnZeroTimeWhenFileHasNoExif(t *testing.T) {
	type args struct {
		testFileName string
	}

	tests := []args{
		{testFileName: testconst.Webp},
		{testFileName: testconst.Png},
	}

	for _, arg := range tests {
		t.Run(arg.testFileName, func(t *testing.T) {
			testFilePath := file.DomainFileTestdataPath(arg.testFileName)

			// 读取测试文件内容
			fileContent := file.NewTestFileContent(t, testFilePath)

			// 提取EXIF创建时间
			actualTime, err := extractExifCreateTime(fileContent.Data)

			// 验证结果
			assert.Error(t, err)
			assert.True(t, actualTime.IsZero(), "EXIF创建时间应该是零时间，但实际是 %v", actualTime.Format("2006-01-02 15:04:05"))
		})
	}
}
