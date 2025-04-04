package file

import (
	"github.com/sirupsen/logrus"
	"time"
)

type MetadataBuilder struct {
	metadata *Metadata
}

func NewMetadataBuilder() *MetadataBuilder {
	return &MetadataBuilder{
		metadata: &Metadata{},
	}
}

func (b *MetadataBuilder) FileName(fileName string) *MetadataBuilder {
	b.metadata.FileName = fileName
	return b
}

func (b *MetadataBuilder) StorageDir(storageDir string) *MetadataBuilder {
	b.metadata.StorageDir = storageDir
	return b
}

func (b *MetadataBuilder) Source(source uint8) *MetadataBuilder {
	b.metadata.Source = source
	return b
}

func (b *MetadataBuilder) ModifiedTime(modifiedTime time.Time) *MetadataBuilder {
	b.metadata.ModifiedTime = modifiedTime
	return b
}

func (b *MetadataBuilder) Build() *Metadata {
	err := b.metadata.validate()
	if err != nil {
		logrus.Errorf("MetadataBuilder.Build: %v", err)
		return nil
	}
	return b.metadata
}
