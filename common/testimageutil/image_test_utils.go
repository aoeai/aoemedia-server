package testimageutil

import (
	"github.com/aoemedia-server/common/testfilecontentutil"
	"github.com/aoemedia-server/domain/file/model"
	imagemodel "github.com/aoemedia-server/domain/image/model"
	"testing"
)

func NewTestAoeImage(t *testing.T, filename string) *imagemodel.AoeImage {
	fileContent := newTestFileContent(t, filename)
	return newTestAoeImage(fileContent)
}

func newTestAoeImage(fileContent *model.FileContent) *imagemodel.AoeImage {
	aoeImage, _ := imagemodel.NewAoeImage(fileContent)
	return aoeImage
}

func newTestFileContent(t *testing.T, filename string) *model.FileContent {
	path := testfilecontentutil.DomainFileModelTestdataPath(filename)
	return testfilecontentutil.NewTestFileContent(t, path)
}
