package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/hobbyqhd/yijing/service/config"
	"github.com/hobbyqhd/yijing/service/middleware"
	"github.com/hobbyqhd/yijing/service/routes"
)

func main() {
	// 初始化配置
	if err := config.Init(); err != nil {
		log.Fatalf("配置初始化失败: %v", err)
	}

	// 创建Gin实例
	r := gin.Default()

	// 使用中间件
	r.Use(middleware.Cors())

	// 注册路由
	routes.RegisterRoutes(r)

	// 启动服务器
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("服务器启动失败: %v", err)
	}
}
