package image

import (
	"errors"
	"testing"
	"time"

	"github.com/aoemedia-server/adapter/driven/persistence/mysql/db"
	"github.com/aoemedia-server/adapter/driven/persistence/mysql/file"
	mysqlimage "github.com/aoemedia-server/adapter/driven/persistence/mysql/image"
	mysqlimagesearch "github.com/aoemedia-server/adapter/driven/persistence/mysql/image_search"
	domainimage "github.com/aoemedia-server/domain/image"
	"gorm.io/gorm"

	"github.com/aoemedia-server/common/testconst"
	"github.com/stretchr/testify/assert"
)

func Test_Upload(t *testing.T) {
	defer teardown(t)

	t.Run("图片上传成功后返回结果正确", shouldReturnCorrectResultAfterImageUploaded)
	t.Run("图片上传成功后file表中存储的数据正确", shouldSaveCorrectDataInFileTableAfterImageUploaded)
	t.Run("图片上传成功后image_upload_record表中存储的数据正确", shouldSaveCorrectDataInImageUploadRecordTableAfterImageUploaded)
	t.Run("图片上传成功后image_search表中存储的数据正确", shouldSaveCorrectDataInImageSearchTableAfterImageUploaded)
}

func shouldReturnCorrectResultAfterImageUploaded(t *testing.T) {
	// 准备测试数据
	domainImage := domainimage.NewTestImage(t, testconst.Jpg)
	userId := int64(1)

	// 执行上传操作
	result, err := Inst().Upload(domainImage, userId)

	// 验证结果
	assert.NoError(t, err, "上传图片时发生错误")
	assert.NotEmpty(t, result.FileId, "文件ID不应为空")
	assert.NotEmpty(t, result.ImageUploadRecordId, "图片上传记录ID不应为空")
	assert.NotEmpty(t, result.FullStoragePath, "存储路径不应为空")
}

func shouldSaveCorrectDataInFileTableAfterImageUploaded(t *testing.T) {
	// 准备测试数据
	domainImage := domainimage.NewTestImage(t, testconst.Png)
	userId := int64(1)

	// 执行上传操作
	result, err := Inst().Upload(domainImage, userId)

	// 验证结果
	assert.NoError(t, err, "上传图片时发生错误")

	// 验证文件表中的数据
	var savedFile file.File
	dbResult := db.Inst().First(&savedFile, result.FileId)
	assert.NoError(t, dbResult.Error, "查询保存的文件记录失败")

	// 验证各个字段是否正确保存
	assert.Equal(t, domainImage.HashValue, savedFile.Hash, "文件哈希值不匹配")
	assert.Equal(t, domainImage.SizeInBytes, savedFile.SizeInBytes, "文件大小不匹配")
	assert.Equal(t, domainImage.FileName, savedFile.Filename, "文件名不匹配")
	assert.Equal(t, domainImage.StorageDir, savedFile.StorageDir, "存储路径不匹配")
}

func shouldSaveCorrectDataInImageUploadRecordTableAfterImageUploaded(t *testing.T) {
	// 准备测试数据
	domainImage := domainimage.NewTestImage(t, testconst.Webp)
	userId := int64(10)

	// 执行上传操作
	result, err := Inst().Upload(domainImage, userId)

	// 验证结果
	assert.NoError(t, err, "上传图片时发生错误")

	// 验证图片上传记录表中的数据
	var savedRecord mysqlimage.ImageUploadRecord
	dbResult := db.Inst().First(&savedRecord, result.ImageUploadRecordId)
	assert.NoError(t, dbResult.Error, "查询保存的图片上传记录失败")

	// 验证各个字段是否正确保存
	assert.Equal(t, userId, savedRecord.UserId, "用户ID不匹配")
	assert.Equal(t, result.FileId, savedRecord.FileId, "文件ID不匹配")
}

func shouldSaveCorrectDataInImageSearchTableAfterImageUploaded(t *testing.T) {
	// 准备测试数据
	domainImage := domainimage.NewTestImage(t, testconst.Pagoda)
	userId := int64(1)

	// 执行上传操作
	result, err := Inst().Upload(domainImage, userId)

	assertImageSearch(t, domainImage, userId, result, err)
}

func assertImageSearch(t *testing.T, domainImage *domainimage.DomainImage, userId int64, result *domainimage.UploadResult, err error) {
	// 验证结果
	assert.NoError(t, err, "上传图片时发生错误")

	// 验证image_search表中的数据
	var savedRecord mysqlimagesearch.ImageSearch
	var dbResult *gorm.DB
	for i := 0; i < 3; i++ {
		dbResult = db.Inst().First(&savedRecord, "file_id = ?", result.FileId)
		if !errors.Is(dbResult.Error, gorm.ErrRecordNotFound) {
			break
		}
		time.Sleep(10 * time.Millisecond)
	}

	assert.NotNil(t, dbResult, "查询保存的图片搜索记录失败")
	assert.NoError(t, dbResult.Error, "查询保存的图片搜索记录失败")

	// 验证各个字段是否正确保存
	assert.Equal(t, userId, savedRecord.UserId, "用户ID不匹配")
	assert.Equal(t, result.FileId, savedRecord.FileId, "文件ID不匹配")
	assert.Equal(t, domainImage.Source, savedRecord.Source, "来源不匹配")
	assert.Equal(t, domainImage.ModifiedTime.Truncate(time.Millisecond), savedRecord.ModifiedTime.Truncate(time.Millisecond), "修改时间不匹配")
	assert.Equal(t, result.FullStoragePath, savedRecord.FullPath, "文件路径不匹配")
	assert.Equal(t, int16(domainImage.ModifiedTime.Year()), savedRecord.Year, "年份不匹配")
	assert.Equal(t, uint8(domainImage.ModifiedTime.Month()), savedRecord.Month, "月份不匹配")
	assert.Equal(t, uint8(domainImage.ModifiedTime.Day()), savedRecord.Day, "日期不匹配")
	assert.False(t, savedRecord.CreatedAt.IsZero(), "创建时间不应为空")
}
