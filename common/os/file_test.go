package aoeos

import (
	"github.com/aoemedia-server/common/testconst"
	aoeos "github.com/aoemedia-server/common/testpath"
	"github.com/aoemedia-server/domain/file"
	"github.com/stretchr/testify/assert"
	"os"
	"path/filepath"
	"testing"
	"time"
)

func TestModTime(t *testing.T) {
	SetupTestFileModTime(t)

	t.Run("当文件是 Txt 时返回正确的修改时间", shouldReturnCorrectModTimeForTxt)
	t.Run("当文件是 Webp 时返回正确的修改时间", shouldReturnCorrectModTimeForWebp)
	t.Run("当文件是 Jpg 时返回正确的修改时间", shouldReturnCorrectModTimeForJpg)
}

// SetupTestFileModTime 设置测试文件的修改时间
func SetupTestFileModTime(t *testing.T) {
	setFileModTime(t, testconst.Txt, time.Date(2025, 2, 8, 7, 40, 25, 362333437, time.UTC))
	setFileModTime(t, testconst.Webp, time.Date(2025, 1, 25, 11, 02, 28, 133320471, time.UTC))
	setFileModTime(t, testconst.Jpg, time.Date(2024, 5, 15, 1, 1, 1, 0, time.UTC))
}

func setFileModTime(t *testing.T, filename string, modTime time.Time) {
	filePath := file.DomainFileTestdataPath(filename)
	err := os.Chtimes(filePath, modTime, modTime)
	if err != nil {
		t.Fatalf("设置文件 %s 的修改时间失败: %v", filepath.Base(filePath), err)
	}
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
	projectRoot, _ := aoeos.ProjectRoot()
	return filepath.Join(projectRoot, "domain", "file", "testdata", filename)
}

func assertModTime(t *testing.T, err error, expectedTime time.Time, modTime time.Time) {
	assert.NoError(t, err)
	assert.Equal(t, expectedTime, modTime, "文件修改时间应该是 %v，但实际是 %v",
		expectedTime.Format("2006-01-02 15:04:05"), modTime.Format("2006-01-02 15:04:05"))
}
