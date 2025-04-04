package testcleanutil

import (
	"github.com/aoemedia-server/domain/file"
	"testing"
)

func CleanTestTempDir(t *testing.T, tempDir string) {
	file.CleanTestTempDir(tempDir)
}

// DeleteTestTempDir 删除测试临时目录
func DeleteTestTempDir() {
	file.DeleteTestTempDir()
}
