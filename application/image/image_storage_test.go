package image

import (
	"github.com/aoemedia-server/adapter/driven/persistence/mysql/db"
	file2 "github.com/aoemedia-server/adapter/driven/persistence/mysql/file"
	file3 "github.com/aoemedia-server/adapter/driven/repository/file"
	"github.com/aoemedia-server/common/testconst"
	"github.com/aoemedia-server/config"
	"github.com/aoemedia-server/domain/file"
	imagemodel "github.com/aoemedia-server/domain/image"
	"github.com/stretchr/testify/assert"
	"path/filepath"
	"testing"
	"time"
)

func Test_createTimeOf(t *testing.T) {
	type args struct {
		name               string
		image              *imagemodel.DomainImage
		expectedPathSuffix string
	}

	nowYearMonth := YearMonthOf(time.Now())

	tests := []args{
		{"当图片的Exif中提取创建时间成功时，使用创建时间的「年-月」做文件夹名",
			imagemodel.NewTestImage(t, testconst.Jpg), filepath.Join("2024-05")},
		{"当图片的Exif中提取创建时间失败时，使用当前时间的「年-月」做文件夹名",
			imagemodel.NewTestImage(t, testconst.Webp), filepath.Join(nowYearMonth)},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			pathSuffix := createTimeOf(test.image)

			assert.Equal(t, test.expectedPathSuffix, pathSuffix)
		})
	}
}

func Test_Save(t *testing.T) {
	db.InitTestDB()
	defer file.CleanTestTempDir(t, config.Inst().Storage.ImageRootDir)

	type args struct {
		name         string
		filename     string
		image        *imagemodel.DomainImage
		expectedPath string
	}

	nowYearMonth := YearMonthOf(time.Now())

	tests := []args{
		{"当图片的Exif中提取创建时间成功时，使用创建时间的「年-月」文件夹存储", testconst.Jpg,
			imagemodel.NewTestImage(t, testconst.Jpg),
			filepath.Join(config.Inst().Storage.ImageRootDir, "2024-05", testconst.Jpg)},
		{"当图片的Exif中提取创建时间失败时，使用当前时间的「年-月」文件夹存储", testconst.Webp,
			imagemodel.NewTestImage(t, testconst.Webp),
			filepath.Join(config.Inst().Storage.ImageRootDir, nowYearMonth, testconst.Webp)},
	}

	for _, test := range tests {
		var id int64
		t.Run(test.name, func(t *testing.T) {
			storage, _ := NewImageStorage(test.image, file3.NewRepository())
			imageId, fullPath, _ := storage.Save(test.filename)
			id = imageId

			assert.Equal(t, test.expectedPath, fullPath)
		})

		t.Cleanup(func() { file2.DeleteTestFile(id) })
	}
}
