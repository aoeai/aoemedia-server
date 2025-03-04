package file

type Repository interface {
	Save(file *DomainFile) (int64, error)
}
