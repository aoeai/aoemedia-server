package file

import (
	"github.com/aoemedia-server/common/testconst"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFileContent_SizeInBytes(t *testing.T) {
	t.Run("当文件是 txt 时返回正确的字节大小", shouldReturnCorrectBytesSizeForTextFile)
	t.Run("当文件是 jpg 时返回正确的字节大小", shouldReturnCorrectBytesSizeForJpgFile)
}

func shouldReturnCorrectBytesSizeForTextFile(t *testing.T) {
	fileContent := newTestTxtFileContent(t)
	assertFileSize(t, fileContent, int64(12))
}

func shouldReturnCorrectBytesSizeForJpgFile(t *testing.T) {
	fileContent := newTestJpgContent(t)
	assertFileSize(t, fileContent, int64(2835185))
}

func TestFileContent_Hash(t *testing.T) {
	t.Run("当文件是 txt 时返回正确的哈希值", shouldReturnCorrectHashValueForTextFile)
	t.Run("当文件是 jpg 时返回正确的哈希值", shouldReturnCorrectHashValueForJpgFile)
}

func shouldReturnCorrectHashValueForTextFile(t *testing.T) {
	fileContent := newTestTxtFileContent(t)
	expectedHash := "76833e55e5f14ac84a134f566e9ff1449adbc0fdbc7e34f3e777688f2f37649c"

	assertHashValue(t, fileContent, expectedHash)
}

func shouldReturnCorrectHashValueForJpgFile(t *testing.T) {
	fileContent := newTestJpgContent(t)
	expectedHash := "f4834082fb18222c0e9704ba04a350d73a87c69d9c794dabf20834f95b194b9b"

	assertHashValue(t, fileContent, expectedHash)
}

func newTestTxtFileContent(t *testing.T) *Content {
	path := DomainFileTestdataPath(testconst.Txt)
	return NewTestFileContent(t, path)
}

func newTestJpgContent(t *testing.T) *Content {
	return NewTestFileContent(t, DomainFileTestdataPath("IMG_20240515_085904.jpg"))
}

func assertFileSize(t *testing.T, fileContent *Content, expectedSize int64) {
	got := fileContent.SizeInBytes()
	assert.Equal(t, expectedSize, got, "文件内容的大小应该是 %v 字节，但实际是 %v 字节", expectedSize, got)
}

func assertHashValue(t *testing.T, fileContent *Content, expectedHash string) {
	got := fileContent.Hash()
	assert.Equal(t, expectedHash, got, "文件内容的哈希值应该是 %v，但实际是 %v", expectedHash, got)
}
