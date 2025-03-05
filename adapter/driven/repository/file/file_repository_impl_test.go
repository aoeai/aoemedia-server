package file

import (
	"github.com/aoemedia-server/adapter/driven/persistence/mysql/db"
	file2 "github.com/aoemedia-server/adapter/driven/persistence/mysql/file"
	"github.com/aoemedia-server/common/testconst"
	domainFile "github.com/aoemedia-server/domain/file"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestSave(t *testing.T) {
	t.Run("当保存成功时应该返回ID", shouldReturnIdWhenSavedSuccess)
}

func shouldReturnIdWhenSavedSuccess(t *testing.T) {
	file := newDomainFile(t)
	id, err := save(file)

	assert.NoError(t, err, "保存文件时不应该返回错误")
	assertSavedSuccess(t, id, file)

	teardown(t, id)
}

func newDomainFile(t *testing.T) *domainFile.DomainFile {
	path := domainFile.DomainFileTestdataPath(testconst.Txt)
	content := domainFile.NewTestFileContent(t, path)
	metadata := domainFile.NewMetadataBuilder().FileName("test.txt").StorageDir(path).Source(1).
		ModifiedTime(time.Date(2025, 3, 2, 22, 31, 15, 0, time.Local)).Build()

	file, _ := domainFile.NewDomainFile(content, metadata)
	return file
}

func assertSavedSuccess(t *testing.T, id int64, file *domainFile.DomainFile) {
	// 验证保存结果
	assert.Greater(t, id, int64(0), "保存文件后应该返回有效的ID")

	// 验证文件内容是否正确保存
	var savedFile file2.File
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
		db.Inst().Delete(&file2.File{}, id)
	})
}
