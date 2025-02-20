package testcleanutil

import (
	"github.com/aoemedia-server/domain/file/storage"
	"testing"
)

func CleanTestTempDir(t *testing.T, tempDir string) {
	storage.CleanTestTempDir(t, tempDir)
}
