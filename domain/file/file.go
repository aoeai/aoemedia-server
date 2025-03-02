package file

type File struct {
	// 文件内容
	content *Content
	// 文件元数据
	metadata *Metadata
}

// Content 获取文件内容
func (f *File) Content() *Content {
	return f.content
}

// Metadata 获取文件元数据
func (f *File) Metadata() *Metadata {
	return f.metadata
}
