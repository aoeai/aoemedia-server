package file

type Storage interface {
	// Save 保存文件到存储系统
	//
	// Returns:
	//   - fullStoragePath 文件的完整存储路径
	//   - err 保存过程中可能发生的错误
	Save(domainFile *DomainFile) (fullStoragePath string, err error)
}
