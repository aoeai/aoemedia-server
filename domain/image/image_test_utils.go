package image

import (
	"github.com/aoemedia-server/domain/file"
	"testing"
)

func NewTestImage(t *testing.T, filename string) *AoeImage {
	fileContent := newTestFileContent(t, filename)
	return newTestImage(fileContent)
}

func newTestImage(fileContent *file.Content) *AoeImage {
	aoeImage, _ := NewAoeImage(fileContent)
	return aoeImage
}

func newTestFileContent(t *testing.T, filename string) *file.Content {
	path := file.DomainFileTestdataPath(filename)
	return file.NewTestFileContent(t, path)
}
