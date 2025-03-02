package file

import "github.com/aoemedia-server/domain/file"

type db struct {
}

func (db *db) Save(file *file.File) *file.File {
	return file
}
