package file

type DomainFile struct {
	// 文件内容
	content *Content
	// 文件元数据
	metadata *Metadata
}

func NewDomainFile(content *Content, metadata *Metadata) (*DomainFile, error) {
	file := &DomainFile{
		content:  content,
		metadata: metadata,
	}

	if err := file.validate(); err != nil {
		return nil, err
	}

	return file, nil
}

// Content 获取文件内容
func (f *DomainFile) Content() *Content {
	return f.content
}

// Metadata 获取文件元数据
func (f *DomainFile) Metadata() *Metadata {
	return f.metadata
}
