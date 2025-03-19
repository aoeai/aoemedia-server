package image

import (
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/aoemedia-server/adapter/driven/persistence/mysql/db"
	mysqlfile "github.com/aoemedia-server/adapter/driven/persistence/mysql/file"
	domainFile "github.com/aoemedia-server/domain/file"

	mysqlimage "github.com/aoemedia-server/adapter/driven/persistence/mysql/image"
	"github.com/aoemedia-server/common/testconst"
	"github.com/aoemedia-server/domain/image"
	domainimage "github.com/aoemedia-server/domain/image"
	"github.com/stretchr/testify/assert"
)

func Test_createTimeOf(t *testing.T) {
	type args struct {
		name               string
		image              *image.DomainImage
		expectedPathSuffix string
	}

	nowYearMonth := time.Now().Format("2006-01")

	tests := []args{
		{
			name:               "jpg",
			image:              domainimage.NewTestImage(t, testconst.Jpg),
			expectedPathSuffix: filepath.Join("2024-05"),
		},
		{
			name:               "webp",
			image:              domainimage.NewTestImage(t, testconst.Webp),
			expectedPathSuffix: filepath.Join(nowYearMonth),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			pathSuffix := createTimeOf(test.image)

			assert.Equal(t, test.expectedPathSuffix, pathSuffix)
		})
	}
}

func TestRepository_Upload(t *testing.T) {
	t.Run("图片上传成功后", func(t *testing.T) {
		t.Run("新上传的文件与原文件相同", newlyUploadedFileIsTheSameAsTheOriginalFile)
		t.Run("file 表中存储的数据正确", dataStoredInFileTableIsCorrect)
		t.Run("image_upload_record 表中存储的数据正确", dataStoredInImageUploadRecordTableIsCorrect)
	})
}

func newlyUploadedFileIsTheSameAsTheOriginalFile(t *testing.T) {
	td := prepareTestUploadData(t)
	defer teardown(t, td)

	storedContent, err := os.ReadFile(td.result.FullStoragePath)
	assert.NoError(t, err, "读取存储的文件失败")
	assert.Equal(t, td.domainImage.Data, storedContent, "存储的文件内容不正确")
}

func dataStoredInFileTableIsCorrect(t *testing.T) {
	td := prepareTestUploadData(t)
	defer teardown(t, td)

	fileRecord, err := getFileById(td.result.FileId)
	assert.NoError(t, err, "获取文件记录失败")
	assert.Equal(t, td.domainImage.HashValue, fileRecord.Hash, "文件哈希值不正确")
	assert.Equal(t, td.domainImage.SizeInBytes, fileRecord.SizeInBytes, "文件大小不正确")
	assert.Equal(t, td.domainImage.FileName, fileRecord.Filename, "文件名不正确")
	assert.Equal(t, td.domainImage.StorageDir, fileRecord.StorageDir, "存储目录不正确")
	assert.Equal(t, td.domainImage.Source, fileRecord.Source, "文件来源不正确")
	assert.Equal(t, td.domainImage.ModifiedTime.Unix(), fileRecord.ModifiedTime.Unix(), "文件修改时间不正确")
}

func dataStoredInImageUploadRecordTableIsCorrect(t *testing.T) {
	td := prepareTestUploadData(t)
	defer teardown(t, td)

	imageUploadRecord, err := getImageUploadRecordById(td.result.ImageUploadRecordId)
	assert.NoError(t, err, "获取图片上传记录失败")
	assert.Equal(t, td.userId, imageUploadRecord.UserId, "用户ID不正确")
	assert.Equal(t, td.result.FileId, imageUploadRecord.FileId, "文件ID不正确")
}

// 测试数据结构体，用于在测试间共享数据
type testUploadData struct {
	domainImage *domainimage.DomainImage
	userId      int64
	result      *image.UploadResult
}

// 准备测试数据
func prepareTestUploadData(t *testing.T) *testUploadData {
	domainImage := domainimage.NewTestImage(t, testconst.Jpg)
	domainImage.Source = 1
	userId := int64(1)

	result, err := Inst().Upload(domainImage, userId)
	assert.NoError(t, err, "上传图片失败")

	return &testUploadData{
		domainImage: domainImage,
		userId:      userId,
		result:      result,
	}
}

// 清理测试数据
func teardown(t *testing.T, td *testUploadData) {
	mysqlfile.DeleteTestFile(td.result.FileId)
	mysqlimage.DeleteTestImageUploadRecordByFileId(td.result.FileId)
	domainFile.CleanTestTempDir(t, td.result.FullStoragePath)
}

func getFileById(id int64) (mysqlfile.File, error) {
	var file mysqlfile.File
	result := db.Inst().Where("id = ?", id).Take(&file)
	return file, result.Error
}

func getImageUploadRecordById(id int64) (mysqlimage.ImageUploadRecord, error) {
	var record mysqlimage.ImageUploadRecord
	result := db.Inst().Where("id = ?", id).Take(&record)
	return record, result.Error
}
