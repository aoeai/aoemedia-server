package image_search

import (
	domainimage "github.com/aoemedia-server/domain/image"
	"github.com/stretchr/testify/mock"
)

type MockImageSearchRepository struct {
	mock.Mock
}

func (m *MockImageSearchRepository) SubscribeImageUploadedEvent(event domainimage.ImageUploadedEvent) {
	m.Called(event)
}

func (m *MockImageSearchRepository) Save(imageSearch ImageSearch) (id int64, error error) {
	args := m.Called(imageSearch)
	return args.Get(0).(int64), args.Error(1)
}

func (m *MockImageSearchRepository) ImageList(params ImageSearchParams) []ImageSearchResult {
	args := m.Called(params)
	return args.Get(0).([]ImageSearchResult)
}
