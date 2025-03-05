package file

import (
	"github.com/aoemedia-server/common/testconst"
	"github.com/aoemedia-server/domain/file"
	"github.com/stretchr/testify/assert"
	"os"
	"path/filepath"
	"testing"
	"time"
)

var testStorageFileRootDir = filepath.Join(os.TempDir(), "temp_txt_success_test")

func TestLocalFileStorage_Store(t *testing.T) {
	t.Run("存储文件成功", shouldStoreFileSuccessfully)
	t.Run("文件名重复时会返回错误", shouldReturnErrorWhenFilenameIsRepeated)
	t.Run("文件存储成功后，存储文件的修改时间与创建时间等于文件元数据中的修改时间", shouldStoreFileWithCorrectModifiedTime)
}

func shouldStoreFileSuccessfully(t *testing.T) {
	defer file.CleanTestTempDir(t, testStorageFileRootDir)

	localStorage := NewLocalFileStorage()
	domainFile := newTestDomainFile(t, testconst.Txt)

	fullStorageDir, err := localStorage.Save(domainFile)

	assert.NoError(t, err, "存储文件失败")
	assert.NotEmpty(t, fullStorageDir, "存储文件后应该返回存储路径")

	expectedFullStorageDir := testStorageFileRootDir
	assert.Equal(t, expectedFullStorageDir, fullStorageDir, "存储文件后返回的存储路径不正确")

	// 验证文件内容
	fullStoragePath := filepath.Join(fullStorageDir, domainFile.FileName)
	storedContent, err := os.ReadFile(fullStoragePath)
	assert.NoError(t, err, "读取存储的文件失败")

	assert.Equal(t, string(domainFile.Data), string(storedContent),
		"存储的文件内容不正确:\n期望的内容: %s\n实际的内容: %s", string(domainFile.Data), string(storedContent))
}

func shouldReturnErrorWhenFilenameIsRepeated(t *testing.T) {
	defer file.CleanTestTempDir(t, testStorageFileRootDir)

	localStorage := NewLocalFileStorage()
	domainFile := newTestDomainFile(t, testconst.Txt)

	// 第一次存储文件
	_, err := localStorage.Save(domainFile)

	assert.NoError(t, err, "第一次存储文件失败")

	// 尝试再次存储同名文件
	_, err = localStorage.Save(domainFile)
	assert.Error(t, err, "存储已存在的文件应该返回错误")
	assert.Contains(t, err.Error(), "文件已经存在", "错误信息应该包含'文件已经存在'")
}

func shouldStoreFileWithCorrectModifiedTime(t *testing.T) {
	defer file.CleanTestTempDir(t, testStorageFileRootDir)

	localStorage := NewLocalFileStorage()
	domainFile := newTestDomainFile(t, testconst.Txt)

	fullStorageDir, err := localStorage.Save(domainFile)
	assert.NoError(t, err, "存储文件失败")

	// 验证文件修改时间
	fullStoragePath := filepath.Join(fullStorageDir, domainFile.FileName)
	fileInfo, err := os.Stat(fullStoragePath)
	assert.NoError(t, err, "获取文件信息失败")

	expectedModTime := domainFile.ModifiedTime
	actualModTime := fileInfo.ModTime()

	assert.Equal(t, expectedModTime.Unix(), actualModTime.Unix(),
		"文件的修改时间应该是 %v，但实际是 %v", expectedModTime, actualModTime)
}

func newTestDomainFile(t *testing.T, filename string) *file.DomainFile {
	content := newTestFileContent(t, filename)
	metadata := newTestFileMetadata(filename)
	domainFile, _ := file.NewDomainFile(content, metadata)
	return domainFile
}

func newTestFileContent(t *testing.T, filename string) *file.Content {
	return file.NewTestFileContent(t, file.DomainFileTestdataPath(filename))
}

func newTestFileMetadata(filename string) *file.Metadata {
	return file.NewMetadataBuilder().FileName(filename).StorageDir(testStorageFileRootDir).Source(1).
		ModifiedTime(time.Date(2025, 3, 2, 22, 31, 15, 0, time.Local)).Build()
}
