package image

import (
	"github.com/aoemedia-server/adapter/driven/persistence/mysql/db"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	// 全局初始化（如启动数据库连接池）
	setup()

	// 运行所有测试
	code := m.Run()

	// 全局清理（如关闭数据库连接）
	//teardown()

	// 退出码传递给 os.Exit
	os.Exit(code)
}

func setup() {
	db.InitTestDB()
}
