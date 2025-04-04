package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConfig_RootDirPath(t *testing.T) {
	t.Run("获取文件存储目录成功", shouldGetFileDirSuccessfully)
}

func shouldGetFileDirSuccessfully(t *testing.T) {
	// 验证RootDirPath返回值
	expectedPath := "temp-test-files"
	actualPath := Inst().StorageFileRootDir()

	assert.Equal(t, expectedPath, actualPath, "获取文件存储目录失败")
}

func TestConfig_UserList(t *testing.T) {
	// 获取用户列表
	users := Inst().UserList()

	// 验证用户列表
	assert.Len(t, users, 2, "用户列表长度应该为2")

	// 验证第一个用户
	assert.Equal(t, "a", users[0].Token, "第一个用户的token应该为'a'")
	assert.Equal(t, int64(1), users[0].ID, "第一个用户的ID应该为1")

	// 验证第二个用户
	assert.Equal(t, "b", users[1].Token, "第二个用户的token应该为'b'")
	assert.Equal(t, int64(2), users[1].ID, "第二个用户的ID应该为2")
}
