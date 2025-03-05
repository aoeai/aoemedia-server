package config

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"os"
	"path/filepath"
	"sync"

	"github.com/BurntSushi/toml"
)

type Config struct {
	// Storage 文件存储相关配置
	Storage StorageConfig `toml:"storage"`
	// Database 数据库相关配置
	Database DatabaseConfig `toml:"database"`
}

type StorageConfig struct {
	FileRootDir  string `toml:"file_root_dir"`
	ImageRootDir string `toml:"image_root_dir"`
}

type DatabaseConfig struct {
	Driver          string `toml:"driver"`
	Host            string `toml:"host"`
	Port            int    `toml:"port"`
	Username        string `toml:"username"`
	Password        string `toml:"password"`
	DBName          string `toml:"dbname"`
	Charset         string `toml:"charset"`
	MaxIdleConns    int    `toml:"max_idle_conns"`
	MaxOpenConns    int    `toml:"max_open_conns"`
	ConnMaxLifetime string `toml:"conn_max_lifetime"`
}

var (
	globalConfig *Config
	initError    error
	once         sync.Once
)

// Inst 获取全局配置，使用sync.Once确保配置只加载一次
func Inst() *Config {
	once.Do(func() {
		globalConfig, initError = loadConfig()
		if initError == nil {
			logrus.Infof("配置加载成功: %+v", globalConfig)
		}
	})

	if initError != nil {
		panic(fmt.Sprintf("配置加载失败: %v", initError))
	}
	return globalConfig
}

func (c *Config) StorageFileRootDir() string {
	return c.Storage.FileRootDir
}

func configFileName() string {
	env := os.Getenv("APP_ENV")
	switch env {
	case "dev":
		return "config.dev.toml"
	case "prod":
		return "config.prod.toml"
	default:
		return "config.test.toml"
	}
}

// loadConfig 从TOML配置文件中加载配置
func loadConfig() (*Config, error) {
	// 获取配置文件路径
	configFile := configFileName()

	// 获取当前目录
	currentDir, err := os.Getwd()
	if err != nil {
		return nil, fmt.Errorf("获取当前目录失败: %w", err)
	}

	// 获取项目根目录（向上查找直到找到 go.mod 文件所在的目录）
	projectRoot := currentDir
	for {
		if _, err := os.Stat(filepath.Join(projectRoot, "go.mod")); err == nil {
			break
		}
		parent := filepath.Dir(projectRoot)
		if parent == projectRoot {
			return nil, fmt.Errorf("无法找到项目根目录")
		}
		projectRoot = parent
	}

	// 构建配置文件的完整路径
	configPath := filepath.Join(projectRoot, "config", configFile)

	// 检查配置文件是否存在
	if _, err := os.Stat(configPath); err != nil {
		return nil, fmt.Errorf("配置文件不存在: %s", configPath)
	}

	// 创建默认配置
	config := &Config{}

	// 解析TOML配置文件
	if _, err := toml.DecodeFile(configPath, config); err != nil {
		return nil, fmt.Errorf("解析配置文件失败: %w", err)
	}

	return config, nil
}
