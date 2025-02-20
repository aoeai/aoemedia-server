package testfilecontentutil

import (
	"github.com/aoemedia-server/domain/file/model"
	"testing"
)

func NewTestFileContent(t *testing.T, filePath string) *model.FileContent {
	return model.NewTestFileContent(t, filePath)
}
