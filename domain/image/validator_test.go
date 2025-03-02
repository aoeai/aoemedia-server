package image

import (
	"github.com/aoemedia-server/common/testconst"
	"github.com/aoemedia-server/domain/file"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsImage(t *testing.T) {
	t.Run("当文件是 txt 时，返回 false", shouldReturnFalseWhenFileIsTxt)
	t.Run("当文件是 jpg 时，返回 true", shouldReturnTrueWhenFileIsJpg)
	t.Run("当文件是 png 时，返回 true", shouldReturnTrueWhenFileIsPng)
	t.Run("当文件是 webp 时，返回 true", shouldReturnTrueWhenFileIsWebp)
}

func shouldReturnFalseWhenFileIsTxt(t *testing.T) {
	assertIsImage(t, testdataPath(testconst.Txt), false)
}

func shouldReturnTrueWhenFileIsJpg(t *testing.T) {
	assertIsImage(t, testdataPath(testconst.Jpg), true)
}

func shouldReturnTrueWhenFileIsPng(t *testing.T) {
	assertIsImage(t, testdataPath(testconst.Png), true)
}

func shouldReturnTrueWhenFileIsWebp(t *testing.T) {
	assertIsImage(t, testdataPath(testconst.Webp), true)
}

func testdataPath(filename string) string {
	return file.DomainFileTestdataPath(filename)
}

func assertIsImage(t *testing.T, filePath string, expected bool) {
	fileContent := file.NewTestFileContent(t, filePath)
	got := IsImage(fileContent)

	assert.Equal(t, expected, got)
}
