package file

import "time"

type MetadataBuilder struct {
	metadata *Metadata
}

func NewMetadataBuilder() *MetadataBuilder {
	return &MetadataBuilder{
		metadata: &Metadata{},
	}
}

func (b *MetadataBuilder) FileName(fileName string) *MetadataBuilder {
	b.metadata.fileName = fileName
	return b
}

func (b *MetadataBuilder) StoragePath(storagePath string) *MetadataBuilder {
	b.metadata.storagePath = storagePath
	return b
}

func (b *MetadataBuilder) Source(source int) *MetadataBuilder {
	b.metadata.source = source
	return b
}

func (b *MetadataBuilder) ModifiedTime(modifiedTime time.Time) *MetadataBuilder {
	b.metadata.modifiedTime = modifiedTime
	return b
}

func (b *MetadataBuilder) Build() *Metadata {
	return b.metadata
}
