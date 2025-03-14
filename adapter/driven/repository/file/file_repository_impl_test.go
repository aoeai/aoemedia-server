package file

import (
	"testing"
	"time"

	"github.com/aoemedia-server/adapter/driven/persistence/mysql/db"
	persistencefile "github.com/aoemedia-server/adapter/driven/persistence/mysql/file"
	"github.com/aoemedia-server/common/testconst"
	domainFile "github.com/aoemedia-server/domain/file"
	"github.com/stretchr/testify/assert"
)

func TestRepository_Save(t *testing.T) {
	t.Run("当保存成功时应该返回ID", shouldReturnIdWhenSaveSuccess)
	t.Run("当文件名和存储目录已存在时应该返回错误", shouldReturnErrorWhenFileExist)
}

func shouldReturnIdWhenSaveSuccess(t *testing.T) {
	// 准备测试数据
	file := newTestDomainFile(t)
	repo := Inst()

	// 执行保存操作
	id, err := repo.Save(file, db.Inst())

	// 验证结果
	assert.NoError(t, err, "保存文件时不应该返回错误")
	assertFileSavedSuccess(t, id, file)

	// 清理测试数据
	teardown(t, id)
}

func shouldReturnErrorWhenFileExist(t *testing.T) {
	// 准备测试数据
	file := newTestDomainFile(t)
	repo := Inst()

	tx := db.Inst()
	// 第一次保存文件
	id, err := repo.Save(file, tx)
	assert.NoError(t, err, "第一次保存文件不应该返回错误")
	defer teardown(t, id)

	// 尝试再次保存相同的文件
	_, err = repo.Save(file, tx)

	// 验证是否返回错误
	assert.Error(t, err, "当文件名和存储目录已存在时应该返回错误")
	assert.Contains(t, err.Error(), "Duplicate entry", "错误信息应该包含数据库唯一约束冲突信息")

	// 清理测试数据
	teardown(t, id)
}

func newTestDomainFile(t *testing.T) *domainFile.DomainFile {
	path := domainFile.DomainFileTestdataPath(testconst.Txt)
	content := domainFile.NewTestFileContent(t, path)
	metadata := domainFile.NewMetadataBuilder().
		FileName("test.txt").
		StorageDir(path).
		Source(1).
		ModifiedTime(time.Date(2025, 3, 2, 22, 31, 15, 0, time.Local)).
		Build()

	file, err := domainFile.NewDomainFile(content, metadata)
	assert.NoError(t, err, "创建测试文件对象失败")
	return file
}

func assertFileSavedSuccess(t *testing.T, id int64, file *domainFile.DomainFile) {
	// 验证保存结果
	assert.Greater(t, id, int64(0), "保存文件后应该返回有效的ID")

	// 验证文件内容是否正确保存
	var savedFile persistencefile.File
	result := db.Inst().First(&savedFile, id)
	assert.NoError(t, result.Error, "查询保存的文件记录失败")

	// 验证各个字段是否正确保存
	assert.Equal(t, file.HashValue, savedFile.Hash, "文件哈希值不匹配")
	assert.Equal(t, file.SizeInBytes, savedFile.SizeInBytes, "文件大小不匹配")
	assert.Equal(t, file.FileName, savedFile.Filename, "文件名不匹配")
	assert.Equal(t, file.StorageDir, savedFile.StorageDir, "存储路径不匹配")
	assert.Equal(t, file.Source, savedFile.Source, "文件来源不匹配")
}

func teardown(t *testing.T, id int64) {
	t.Cleanup(func() {
		db.Inst().Delete(&persistencefile.File{}, id)
	})
}
