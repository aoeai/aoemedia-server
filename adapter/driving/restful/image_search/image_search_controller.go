package image_search

import (
	"github.com/aoemedia-server/adapter/driving/restful/authorization"
	"github.com/aoemedia-server/adapter/driving/restful/response"
	"github.com/aoemedia-server/application/image"
	"github.com/aoemedia-server/common/converter"
	domainimagesearch "github.com/aoemedia-server/domain/image/image_search"
	"github.com/gin-gonic/gin"
	"time"
)

type ImageSearchController struct {
}

func NewImageSearchController() *ImageSearchController {
	return &ImageSearchController{}
}

type ImageSearchRequest struct {
	// 来源 1:相机 2:微信
	Source uint8 `json:"source"`
	// 最早修改时间
	EarliestModifiedTime string `json:"earliest_modified_time"`
}

func (c *ImageSearchController) List(ctx *gin.Context) {
	auth := authorization.NewAuth(ctx)
	if auth.Invalid() {
		response.SendUnauthorized(ctx)
		return
	}

	var json ImageSearchRequest
	if err := ctx.ShouldBindJSON(&json); err != nil {
		response.SendBadRequest(ctx, "无效的请求参数")
		return
	}

	earliestModifiedTime := time.Now()
	if json.EarliestModifiedTime != "" {
		earliestModifiedTime = converter.StrToTime(json.EarliestModifiedTime)
	}
	if earliestModifiedTime.IsZero() {
		response.SendBadRequest(ctx, "earliest_modified_time 格式不正确")
		return
	}

	params := domainimagesearch.ImageSearchParams{
		UserId:       auth.UserId,
		ModifiedTime: earliestModifiedTime,
		Source:       json.Source,
		Limit:        10,
	}
	result := image.NewSearcher().ImageList(params)

	response.SendSuccess(ctx, convertToListResult(result))
}
