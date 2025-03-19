package image_search

import (
	"github.com/aoemedia-server/adapter/driven/persistence/mysql/db"
	domainimagesearch "github.com/aoemedia-server/domain/image_search"
)

func Create(imageSearch domainimagesearch.ImageSearch) (id int64, error error) {
	entity := newEntity(imageSearch)
	tx := db.Inst().Create(&entity)

	return entity.ID, tx.Error
}

func newEntity(imageSearch domainimagesearch.ImageSearch) ImageSearch {
	entity := ImageSearch{}
	entity.UserId = imageSearch.UserId
	entity.FileId = imageSearch.FileId
	entity.Source = imageSearch.Source
	entity.ModifiedTime = imageSearch.ModifiedTime
	entity.FullPath = imageSearch.FullPath
	entity.Year = imageSearch.Year
	entity.Month = imageSearch.Month
	entity.Day = imageSearch.Day

	return entity
}
