package config

import (
	"os"

	"github.com/joho/godotenv"
)

// LoadEnv 加载环境变量
func LoadEnv() error {
	// 加载.env文件
	if err := godotenv.Load(); err != nil {
		return err
	}
	return nil
}

// GetEnv 获取环境变量的值
func GetEnv(key string) string {
	return os.Getenv(key)
}
