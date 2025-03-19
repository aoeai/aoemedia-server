package image

import (
	"github.com/aoemedia-server/domain/file"
	"testing"
	"time"
)

func NewTestImage(t *testing.T, filename string) *DomainImage {
	fileContent := newTestFileContent(t, filename)
	metadata := file.NewMetadataBuilder().FileName(filename).
		StorageDir(file.DomainFileTestdataDir()).Source(1).
		ModifiedTime(time.Now()).Build()
	domainFile, _ := file.NewDomainFile(fileContent, metadata)

	return newTestImage(domainFile)
}

func newTestImage(domainFile *file.DomainFile) *DomainImage {
	domainImage, _ := New(domainFile)
	return domainImage
}

func newTestFileContent(t *testing.T, filename string) *file.Content {
	path := file.DomainFileTestdataPath(filename)
	return file.NewTestFileContent(t, path)
}
