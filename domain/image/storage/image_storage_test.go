package storage

import (
	"github.com/aoemedia-server/common/testcleanutil"
	"github.com/aoemedia-server/common/testconst"
	"github.com/aoemedia-server/common/testimageutil"
	"github.com/aoemedia-server/config"
	"github.com/stretchr/testify/assert"
	"path/filepath"
	"testing"
)

func TestImageStorage_Save(t *testing.T) {
	t.Run("当从 EXIF 中获取创建时间成功时，文件的修改时间会被设置为 EXIF 中的创建时间", shouldSetFileModificationTimeToExifCreateTimeWhenExtractExifCreateTimeSuccessfully)
}

// 当从 EXIF 中获取创建时间成功时，文件的修改时间会被设置为 EXIF 中的创建时间
func shouldSetFileModificationTimeToExifCreateTimeWhenExtractExifCreateTimeSuccessfully(t *testing.T) {
	dir := config.Instance().FileStorage.ImageDir
	defer testcleanutil.CleanTestTempDir(t, dir)

	filename := testconst.Jpg
	aoeImage := testimageutil.NewTestAoeImage(t, filename)

	imageStorage, _ := NewImageStorage(dir)
	fullPath, err := imageStorage.Save(aoeImage, filename)

	assert.NoError(t, err)

	expectFullPath := filepath.Join(dir, filename)
	assert.Equal(t, expectFullPath, fullPath, "保存的文件名应该是 %v，但实际是 %v", filename, fullPath)
}
