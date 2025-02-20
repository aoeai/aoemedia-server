package storage

import (
	"github.com/aoemedia-server/application/storage/path"
	"github.com/aoemedia-server/common/testcleanutil"
	"github.com/aoemedia-server/common/testconst"
	"github.com/aoemedia-server/common/testimageutil"
	"github.com/aoemedia-server/config"
	imagemodel "github.com/aoemedia-server/domain/image/model"
	"github.com/stretchr/testify/assert"
	"path/filepath"
	"testing"
	"time"
)

func Test_createTimeOf(t *testing.T) {
	type args struct {
		name               string
		image              *imagemodel.AoeImage
		expectedPathSuffix string
	}

	nowYearMonth := path.YearMonthOf(time.Now())

	tests := []args{
		{"当图片的Exif中提取创建时间成功时，使用创建时间的「年-月」做文件夹名",
			testimageutil.NewTestAoeImage(t, testconst.Jpg), filepath.Join("2024-05")},
		{"当图片的Exif中提取创建时间失败时，使用当前时间的「年-月」做文件夹名",
			testimageutil.NewTestAoeImage(t, testconst.Webp), filepath.Join(nowYearMonth)},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			pathSuffix := createTimeOf(test.image)

			assert.Equal(t, test.expectedPathSuffix, pathSuffix)
		})
	}
}

func Test_Save(t *testing.T) {
	defer testcleanutil.CleanTestTempDir(t, config.Instance().FileStorage.ImageDir)

	type args struct {
		name         string
		filename     string
		image        *imagemodel.AoeImage
		expectedPath string
	}

	nowYearMonth := path.YearMonthOf(time.Now())

	tests := []args{
		{"当图片的Exif中提取创建时间成功时，使用创建时间的「年-月」文件夹存储", testconst.Jpg,
			testimageutil.NewTestAoeImage(t, testconst.Jpg),
			filepath.Join(config.Instance().FileStorage.ImageDir, "2024-05", testconst.Jpg)},
		{"当图片的Exif中提取创建时间失败时，使用当前时间的「年-月」文件夹存储", testconst.Webp,
			testimageutil.NewTestAoeImage(t, testconst.Webp),
			filepath.Join(config.Instance().FileStorage.ImageDir, nowYearMonth, testconst.Webp)},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			imageStorage, _ := NewImageStorage(test.image)
			fullPath, _ := imageStorage.Save(test.filename)

			assert.Equal(t, test.expectedPath, fullPath)
		})
	}
}
