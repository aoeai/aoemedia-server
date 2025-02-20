package main

import (
	"fmt"
	"github.com/aoemedia-server/adapter/driving/restful/upload"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"os"
)

func main() {
	if err := setupEnv(); err != nil {
		logrus.Fatalf("环境变量设置失败: %v", err)
		return
	}

	// 初始化控制器
	// 初始化Gin引擎
	r := gin.Default()

	// 配置路由
	r.POST(upload.File, upload.NewFileController().Upload)
	r.POST(upload.Image, upload.NewImageController().Upload)

	// 启动服务器
	err := r.Run(":8080")
	if err != nil {
		logrus.Fatalf("服务器启动失败: %v", err)
	}
}

func setupEnv() error {
	// 如果环境变量未设置，则使用默认值"dev"
	if os.Getenv("APP_ENV") == "" {
		if err := os.Setenv("APP_ENV", "dev"); err != nil {
			return fmt.Errorf("设置环境变量失败: %w", err)
		}
	}
	// 使用log打印环境变量
	logrus.Printf("当前环境变量: %s", os.Getenv("APP_ENV"))
	return nil
}
