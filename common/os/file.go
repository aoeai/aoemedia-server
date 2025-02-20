package aoeos

import (
	"os"
	"time"
)

// ModTime returns the modification time of a file at the specified path.
//
// Parameters:
//   - path: The file path to check
//
// Returns:
//   - modTime: The modification time of the file
//   - err: An error if the file cannot be accessed or does not exist
func ModTime(path string) (modTime time.Time, err error) {
	fileInfo, err := os.Stat(path)
	if err != nil {
		return time.Time{}, err
	}
	return fileInfo.ModTime(), nil
}
