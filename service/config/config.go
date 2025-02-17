package config

import (
	"context"
	"fmt"
	"os"

	"github.com/go-redis/redis/v8"
	gopenai "github.com/sashabaranov/go-openai"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB           *gorm.DB
	RedisClient  *redis.Client
	OpenAIClient *gopenai.Client
)

func Init() error {
	// 加载环境变量
	if err := LoadEnv(); err != nil {
		return fmt.Errorf("环境变量加载失败: %v", err)
	}

	// 初始化MySQL连接
	if err := initMySQL(); err != nil {
		return fmt.Errorf("MySQL初始化失败: %v", err)
	}

	// 初始化Redis连接
	if err := initRedis(); err != nil {
		return fmt.Errorf("Redis初始化失败: %v", err)
	}

	// 初始化OpenAI客户端
	initOpenAI()

	return nil
}

func initMySQL() error {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		os.Getenv("MYSQL_USER"),
		os.Getenv("MYSQL_PASSWORD"),
		os.Getenv("MYSQL_HOST"),
		os.Getenv("MYSQL_PORT"),
		os.Getenv("MYSQL_DATABASE"),
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}

	DB = db
	return nil
}

func initRedis() error {
	RedisClient = redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%s",
			os.Getenv("REDIS_HOST"),
			os.Getenv("REDIS_PORT"),
		),
		Password: os.Getenv("REDIS_PASSWORD"),
		DB:       0,
	})

	// 测试连接
	ctx := context.Background()
	_, err := RedisClient.Ping(ctx).Result()
	return err
}

func initOpenAI() {
	OpenAIClient = gopenai.NewClient(os.Getenv("OPENAI_API_KEY"))
}
