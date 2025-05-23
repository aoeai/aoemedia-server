package file

import (
	"github.com/aoemedia-server/common/testpath"
	"github.com/aoemedia-server/config"
	"github.com/sirupsen/logrus"
	"os"
	"path/filepath"
	"testing"
)

func DomainFileTestdataDir() string {
	projectRoot, _ := testpath.ProjectRoot()
	return filepath.Join(projectRoot, "domain", "file", "testdata")
}

func DomainFileTestdataPath(filename string) string {
	return filepath.Join(DomainFileTestdataDir(), filename)
}

func NewTestFileContent(t *testing.T, filePath string) *Content {
	// 准备测试数据
	data, err := os.ReadFile(filePath)
	if err != nil {
		t.Fatalf("无法读取测试文件: %v", err)
	}

	// 当创建一个新的文件内容对象时
	fileContent := NewFileContent(data)
	return fileContent
}

func CleanTestTempDir(tempDir string) {
	// 删除临时目录及其所有内容
	err := os.RemoveAll(tempDir)
	if err != nil {
		logrus.Errorf("清理临时目录失败: %v", err)
	}
}

// DeleteTestTempDir 删除测试临时目录
func DeleteTestTempDir() {
	err := os.RemoveAll(config.Inst().Storage.ImageRootDir)
	if err != nil {
		logrus.Errorf("CleanTestCurrentDataDir 清理临时目录失败: %v", err)
	}
}
