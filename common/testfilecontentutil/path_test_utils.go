package testfilecontentutil

import (
	"github.com/aoemedia-server/common/testpath"
	"path/filepath"
)

func DomainFileModelTestdataPath(filename string) string {
	projectRoot, _ := testpath.ProjectRoot()
	return filepath.Join(projectRoot, "domain", "file", "model", "testdata", filename)
}
