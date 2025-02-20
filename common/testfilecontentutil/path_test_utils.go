package testfilecontentutil

import (
	aoeos "github.com/aoemedia-server/common/os"
	"path/filepath"
)

func DomainFileModelTestdataPath(filename string) string {
	projectRoot, _ := aoeos.ProjectRoot()
	return filepath.Join(projectRoot, "domain", "file", "model", "testdata", filename)
}
