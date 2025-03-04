package config

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestConfig_RootDirPath(t *testing.T) {
	t.Run("获取文件存储目录成功", shouldGetFileDirSuccessfully)
}

func shouldGetFileDirSuccessfully(t *testing.T) {
	// 获取全局配置实例
	config := Inst()

	// 验证RootDirPath返回值
	expectedPath := ".temp-test-files"
	actualPath := config.RootDirPath()

	assert.Equal(t, expectedPath, actualPath, "获取文件存储目录失败")
}
