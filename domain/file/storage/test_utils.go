package storage

import (
	"os"
	"testing"
)

func CleanTestTempDir(t *testing.T, tempDir string) {
	// 删除临时目录及其所有内容
	err := os.RemoveAll(tempDir)
	if err != nil {
		t.Errorf("清理临时目录失败: %v", err)
	}
}
