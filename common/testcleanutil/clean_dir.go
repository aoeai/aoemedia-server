package testcleanutil

import (
	"github.com/aoemedia-server/domain/file"
	"testing"
)

func CleanTestTempDir(t *testing.T, tempDir string) {
	file.CleanTestTempDir(t, tempDir)
}
