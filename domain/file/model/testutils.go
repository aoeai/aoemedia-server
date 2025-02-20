package model

import (
	aoeos "github.com/aoemedia-server/common/os"
	"os"
	"path/filepath"
	"testing"
)

func DomainFileModelTestdataPath(filename string) string {
	projectRoot, _ := aoeos.ProjectRoot()
	return filepath.Join(projectRoot, "domain", "file", "model", "testdata", filename)
}

func NewTestFileContent(t *testing.T, filePath string) *FileContent {
	// 准备测试数据
	data, err := os.ReadFile(filePath)
	if err != nil {
		t.Fatalf("无法读取测试文件: %v", err)
	}

	// 当创建一个新的文件内容对象时
	fileContent := NewFileContent(data)
	return fileContent
}
