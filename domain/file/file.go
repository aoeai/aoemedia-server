package file

type DomainFile struct {
	// 文件内容
	*Content
	// 文件元数据
	*Metadata
}

func NewDomainFile(content *Content, metadata *Metadata) (*DomainFile, error) {
	file := &DomainFile{
		Content:  content,
		Metadata: metadata,
	}

	if err := file.validate(); err != nil {
		return nil, err
	}

	return file, nil
}
