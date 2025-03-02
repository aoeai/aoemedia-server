package file

type Repository interface {
	Save(file *File) *File
}
