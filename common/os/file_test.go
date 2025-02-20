package aoeos

import (
	"github.com/aoemedia-server/common/testconst"
	"github.com/stretchr/testify/assert"
	"path/filepath"
	"testing"
	"time"
)

func TestModTime(t *testing.T) {
	t.Run("当文件是 Txt 时返回正确的修改时间", shouldReturnCorrectModTimeForTxt)
	t.Run("当文件是 Webp 时返回正确的修改时间", shouldReturnCorrectModTimeForWebp)
	t.Run("当文件是 Jpg 时返回正确的修改时间", shouldReturnCorrectModTimeForJpg)
}

func shouldReturnCorrectModTimeForTxt(t *testing.T) {
	expectedTime := time.Date(2025, 2, 8, 7, 40, 25, 362333437, time.UTC)
	modTime, err := ModTime(testdataPath(testconst.Txt))

	assertModTime(t, err, expectedTime, modTime)
}

func shouldReturnCorrectModTimeForWebp(t *testing.T) {
	expectedTime := time.Date(2025, 1, 25, 11, 02, 28, 133320471, time.UTC)
	modTime, err := ModTime(testdataPath(testconst.Webp))

	assertModTime(t, err, expectedTime, modTime)
}

func shouldReturnCorrectModTimeForJpg(t *testing.T) {
	expectedTime := time.Date(2024, 5, 15, 1, 1, 1, 0, time.UTC)
	modTime, err := ModTime(testdataPath(testconst.Jpg))

	assertModTime(t, err, expectedTime, modTime)
}

func testdataPath(filename string) string {
	return filepath.Join("..", "..", "domain", "file", "model", "testdata", filename)
}

func assertModTime(t *testing.T, err error, expectedTime time.Time, modTime time.Time) {
	assert.NoError(t, err)
	assert.Equal(t, expectedTime, modTime, "文件修改时间应该是 %v，但实际是 %v",
		expectedTime.Format("2006-01-02 15:04:05"), modTime.Format("2006-01-02 15:04:05"))
}
