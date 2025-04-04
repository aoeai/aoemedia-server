package upload

import (
	"github.com/aoemedia-server/adapter/driven/persistence/mysql/db"
	"github.com/aoemedia-server/adapter/driven/persistence/mysql/file"
	mysqlimage "github.com/aoemedia-server/adapter/driven/persistence/mysql/image"
	mysqlimagesearch "github.com/aoemedia-server/adapter/driven/persistence/mysql/image_search"
	"github.com/aoemedia-server/common/testcleanutil"
	"github.com/aoemedia-server/config"
	domainfile "github.com/aoemedia-server/domain/file"
	"github.com/sirupsen/logrus"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	// 全局初始化（如启动数据库连接池）
	setup()

	// 运行所有测试
	code := m.Run()

	// 全局清理（如关闭数据库连接）
	teardown()

	// 退出码传递给 os.Exit
	os.Exit(code)
}

func setup() {
	db.InitTestDB()
}

func teardown() {
	domainfile.CleanTestTempDir(config.Inst().Storage.ImageRootDir)

	// 使用新会话执行删除
	tx := db.InstForceDelete()
	tx.Delete(&file.File{})
	logrus.Infof("teardown: 删除 file %v 条数据", tx.RowsAffected)

	tx = db.InstForceDelete()
	tx.Delete(&mysqlimage.ImageUploadRecord{})
	logrus.Infof("teardown: 删除 image_upload_record %v 条数据", tx.RowsAffected)

	tx = db.InstForceDelete()
	tx.Delete(&mysqlimagesearch.ImageSearch{})
	logrus.Infof("teardown: 删除 image_search %v 条数据", tx.RowsAffected)

	testcleanutil.DeleteTestTempDir()
}
