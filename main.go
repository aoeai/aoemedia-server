package main

import (
	"github.com/aoemedia-server/adapter/driven/persistence/mysql/db"
	"github.com/aoemedia-server/adapter/driving/restful/upload"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"os"
)

func main() {
	initEnv()
	db.InitDB()
	initEngine()
}

func initEnv() {
	// 如果环境变量未设置，则使用默认值"dev"
	if os.Getenv("APP_ENV") == "" {
		if err := os.Setenv("APP_ENV", "dev"); err != nil {
			logrus.Fatalf("环境变量设置失败: %v", err)
		}
	}
	logrus.Printf("当前环境变量: %s", os.Getenv("APP_ENV"))
}

func initEngine() {
	engine := gin.Default()
	setupRoutes(engine)
	startEngine(engine)
}

// setupRoutes 配置路由
func setupRoutes(r *gin.Engine) {
	r.POST(upload.File, upload.NewFileController().Upload)
	r.POST(upload.Image, upload.NewImageController().Upload)
}

func startEngine(ginEngine *gin.Engine) {
	err := ginEngine.Run(":8080")
	if err != nil {
		logrus.Fatalf("服务器启动失败: %v", err)
	}
}
