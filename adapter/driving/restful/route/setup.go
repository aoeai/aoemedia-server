package route

import (
	"github.com/aoemedia-server/adapter/driving/restful/image_search"
	"github.com/aoemedia-server/adapter/driving/restful/route/url"
	"github.com/aoemedia-server/adapter/driving/restful/upload"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func InitEngine() {
	router := gin.Default()
	if setTrustedProxies(router) {
		return
	}
	setupRoutes(router)
	startEngine(router)
}

func setTrustedProxies(router *gin.Engine) bool {
	// 如果你没有使用上述函数指定可信代理，Gin 默认会信任所有代理，这并不安全。
	// 同时，如果你不使用任何代理，可以通过 Engine.SetTrustedProxies(nil) 来禁用此功能，
	// 这样 Context.ClientIP() 将直接返回远程地址，避免不必要的计算。
	// https://gin-gonic.com/zh-cn/docs/deployment/
	err := router.SetTrustedProxies(nil)
	if err != nil {
		logrus.Fatalf("设置受信任代理失败: %v", err)
		return true
	}
	return false
}

func startEngine(ginEngine *gin.Engine) {
	err := ginEngine.Run(":8080")
	if err != nil {
		logrus.Fatalf("服务器启动失败: %v", err)
	}
}

// setupRoutes 配置路由
func setupRoutes(r *gin.Engine) {
	r.POST(url.File, upload.NewFileController().Upload)
	r.POST(url.Image, upload.NewImageController().Upload)

	r.GET(url.ImageSearch, image_search.NewImageSearchController().List)
}
