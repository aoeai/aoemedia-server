package file

import "github.com/aoemedia-server/adapter/driven/persistence/db"

func DeleteTestFile(id int64) {
	db.Inst().Delete(&File{}, id)
}
