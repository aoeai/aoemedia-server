package main

import (
	"github.com/aoemedia-server/adapter/driven/persistence/mysql/db"
	"github.com/aoemedia-server/adapter/driving/restful/route"
	"github.com/sirupsen/logrus"
	"os"
)

func main() {
	initEnv()
	db.InitDB()
	route.InitEngine()
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
