package upload

import (
	"github.com/aoemedia-server/adapter/driven/persistence/db"
	"github.com/aoemedia-server/common/testconst"
	"github.com/aoemedia-server/config"
	"github.com/aoemedia-server/domain/file"
	"testing"
)

func TestImageController_Upload(t *testing.T) {
	file.CleanTestTempDir(t, config.Inst().FileStorage.ImageDir)
	db.InitTestDB()

	t.Run("上传Txt文件会返回错误信息：文件内容不是图片类型", func(t *testing.T) {
		testFilePath := file.DomainFileTestdataPath(testconst.Txt)
		assertBadRequest(t, testFilePath, "文件内容不是图片类型", Image)
	})

	t.Run("上传Jpg图片成功", func(t *testing.T) {
		testFilePath := file.DomainFileTestdataPath(testconst.Jpg)

		code, response := postFile(t, testFilePath, Image)

		assertSuccess(t, code, response, testconst.Jpg,
			"f4834082fb18222c0e9704ba04a350d73a87c69d9c794dabf20834f95b194b9b", float64(2835185))

		deleteImageFileByDB(response)
	})

	t.Run("上传Webp图片成功", func(t *testing.T) {
		testFilePath := file.DomainFileTestdataPath(testconst.Webp)

		code, response := postFile(t, testFilePath, Image)

		assertSuccess(t, code, response, testconst.Webp,
			"548d859e1efa5f6d3d31aa8c444f7028f31bd4054707acbc77bfa20e948aeeb2", float64(98700))

		deleteImageFileByDB(response)
	})

	t.Cleanup(func() { file.CleanTestTempDir(t, config.Inst().FileStorage.ImageDir) })
}
