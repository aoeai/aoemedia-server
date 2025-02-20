package upload

import (
	"github.com/aoemedia-server/common/testcleanutil"
	"github.com/aoemedia-server/common/testconst"
	"github.com/aoemedia-server/common/testfilecontentutil"
	"github.com/aoemedia-server/config"
	"testing"
)

func TestImageController_Upload(t *testing.T) {
	defer testcleanutil.CleanTestTempDir(t, config.Instance().FileStorage.ImageDir)

	t.Run("上传Txt文件会返回错误信息：文件内容不是图片类型", func(t *testing.T) {
		testFilePath := testfilecontentutil.DomainFileModelTestdataPath(testconst.Txt)
		assertBadRequest(t, testFilePath, "文件内容不是图片类型", Image)
	})

	t.Run("上传Jpg图片成功", func(t *testing.T) {
		testFilePath := testfilecontentutil.DomainFileModelTestdataPath(testconst.Jpg)

		assertSuccess(t, testFilePath, testconst.Jpg,
			"f4834082fb18222c0e9704ba04a350d73a87c69d9c794dabf20834f95b194b9b", Image, float64(2835185))
	})

	t.Run("上传Webp图片成功", func(t *testing.T) {
		testFilePath := testfilecontentutil.DomainFileModelTestdataPath(testconst.Webp)

		assertSuccess(t, testFilePath, testconst.Webp,
			"548d859e1efa5f6d3d31aa8c444f7028f31bd4054707acbc77bfa20e948aeeb2", Image, float64(98700))
	})
}
