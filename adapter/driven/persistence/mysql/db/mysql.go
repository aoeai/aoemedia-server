package db

import (
	"fmt"
	"github.com/aoemedia-server/config"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"sync"
	"time"
)

var dbOnce sync.Once

// DB 全局数据库连接实例
var dbInst *gorm.DB

func InitTestDB() {
	InitDB()
}

func InitDB() {
	dbOnce.Do(func() {
		err := initDB(config.Inst())
		if err != nil {
			logrus.Fatalf("数据库连接初始化失败: %v", err)
		}
	})
}

// initDB 初始化数据库连接
func initDB(cfg *config.Config) error {
	// 构建DSN连接字符串
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=True&loc=Local",
		cfg.Database.Username,
		cfg.Database.Password,
		cfg.Database.Host,
		cfg.Database.Port,
		cfg.Database.DBName,
		cfg.Database.Charset,
	)

	// 设置GORM配置
	gormConfig := &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, // 使用单数表名
		},
	}

	// 连接数据库
	db, err := gorm.Open(mysql.Open(dsn), gormConfig)
	if err != nil {
		return fmt.Errorf("连接数据库失败: %v", err)
	}

	// 获取底层的sqlDB
	sqlDB, err := db.DB()
	if err != nil {
		return fmt.Errorf("获取sqlDB实例失败: %v", err)
	}

	// 设置连接池参数
	sqlDB.SetMaxIdleConns(cfg.Database.MaxIdleConns)
	sqlDB.SetMaxOpenConns(cfg.Database.MaxOpenConns)
	sqlDB.SetConnMaxLifetime(time.Hour) // 设置连接最大生命周期

	// 设置全局DB实例
	dbInst = db
	logrus.Infof("数据库连接成功: Host:%v Port:%v DBName:%v Charset:%v",
		cfg.Database.Host, cfg.Database.Port, cfg.Database.DBName, cfg.Database.Charset)
	return nil
}

// Inst 获取数据库连接实例
func Inst() *gorm.DB {
	return dbInst
}

// InstForceDelete 获取数据库连接实例（强制删除）
func InstForceDelete() *gorm.DB {
	return Inst().Session(&gorm.Session{AllowGlobalUpdate: true}).Unscoped()
}
