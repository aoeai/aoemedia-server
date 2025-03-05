package file

import (
	"github.com/aoemedia-server/common/testconst"
	"github.com/stretchr/testify/assert"
	"os"
	"path/filepath"
	"testing"
)

func TestLocalFileStorage_Store(t *testing.T) {
	t.Run("存储文件成功", shouldStoreFileSuccessfully)
	t.Run("文件名重复时会返回错误", shouldReturnErrorWhenFilenameIsRepeated)
}

func shouldStoreFileSuccessfully(t *testing.T) {
	// 准备测试数据
	fileContent := newTestFileContent(t)

	// 创建临时目录作为测试存储路径
	tempDir := filepath.Join(os.TempDir(), "temp_txt_success_test")
	defer CleanTestTempDir(t, tempDir)

	// 创建存储器
	storage, err := NewLocalFileStorage(tempDir)
	assert.NoError(t, err, "创建存储器失败 %s", tempDir)

	// 存储文件
	testFileName := "test.txt"
	relativePath, err := storage.Save(fileContent, testFileName)
	assert.NoError(t, err, "存储文件失败")

	// 验证文件内容
	storePath := filepath.Join(tempDir, relativePath)
	storedContent, err := os.ReadFile(storePath)
	assert.NoError(t, err, "读取存储的文件失败")

	assert.Equal(t, string(fileContent.Data), string(storedContent),
		"存储的文件内容不正确:\n期望的内容: %s\n实际的内容: %s", string(fileContent.Data), string(storedContent))
}

func shouldReturnErrorWhenFilenameIsRepeated(t *testing.T) {
	// 准备测试数据
	fileContent := newTestFileContent(t)

	// 创建临时目录作为测试存储路径
	tempDir := filepath.Join(os.TempDir(), "temp_txt_exists_test")
	defer CleanTestTempDir(t, tempDir)

	// 创建存储器
	storage, err := NewLocalFileStorage(tempDir)
	assert.NoError(t, err, "创建存储器失败 %s", tempDir)

	// 第一次存储文件
	testFileName := "test.txt"
	_, err = storage.Save(fileContent, testFileName)
	assert.NoError(t, err, "第一次存储文件失败")

	// 尝试再次存储同名文件
	_, err = storage.Save(fileContent, testFileName)
	assert.Error(t, err, "存储已存在的文件应该返回错误")
	assert.Contains(t, err.Error(), "文件已经存在", "错误信息应该包含'文件已经存在'")
}

func newTestFileContent(t *testing.T) *Content {
	return NewTestFileContent(t, DomainFileTestdataPath(testconst.Txt))
}
