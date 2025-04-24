package image_search

import (
	domainimagesearch "github.com/aoemedia-server/domain/image/image_search"
	"github.com/stretchr/testify/mock"
)

// MockImageSearchQuery 是ImageSearchQuery接口的mock实现
type MockImageSearchQuery struct {
	mock.Mock
}

func (m *MockImageSearchQuery) ExistByFileId(fileId int64) bool {
	args := m.Called(fileId)
	return args.Bool(0)
}

func (m *MockImageSearchQuery) ImageList(params domainimagesearch.ImageSearchParams) []ImageSearch {
	args := m.Called(params)
	return args.Get(0).([]ImageSearch)
}
