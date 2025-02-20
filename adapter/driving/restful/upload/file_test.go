package upload

import (
	"github.com/aoemedia-server/common/testcleanutil"
	"github.com/aoemedia-server/common/testconst"
	"github.com/aoemedia-server/common/testfilecontentutil"
	"github.com/aoemedia-server/config"
	"testing"
)

func TestFileController_Upload(t *testing.T) {
	defer testcleanutil.CleanTestTempDir(t, config.Instance().RootDirPath())

	t.Run("上传 Txt 文件成功", func(t *testing.T) {
		testFilePath := testfilecontentutil.DomainFileModelTestdataPath(testconst.Txt)
		assertSuccess(t, testFilePath, testconst.Txt,
			"76833e55e5f14ac84a134f566e9ff1449adbc0fdbc7e34f3e777688f2f37649c", File, float64(12))
	})

	t.Run("上传 Jpg 文件成功", func(t *testing.T) {
		testFilePath := testfilecontentutil.DomainFileModelTestdataPath(testconst.Jpg)

		assertSuccess(t, testFilePath, testconst.Jpg,
			"f4834082fb18222c0e9704ba04a350d73a87c69d9c794dabf20834f95b194b9b", File, float64(2835185))
	})
}
