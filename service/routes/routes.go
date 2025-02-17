package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/hobbyqhd/yijing/service/handlers"
	"github.com/hobbyqhd/yijing/service/middleware"
)

func RegisterRoutes(r *gin.Engine) {
	// 用户相关路由
	userGroup := r.Group("/user")
	{
		userHandler := handlers.NewUserHandler()
		userGroup.POST("/register", userHandler.Register)
		userGroup.POST("/login", userHandler.Login)

		// 需要认证的路由
		authorized := userGroup.Use(middleware.Auth())
		authorized.GET("/info", userHandler.GetUserInfo)
		authorized.PUT("/info", userHandler.UpdateUserInfo)
	}

	// 占卜相关路由
	divinationGroup := r.Group("/divination")
	{
		divinationHandler := handlers.NewDivinationHandler()
		// 应用认证中间件
		authorized := divinationGroup.Use(middleware.Auth())
		authorized.POST("", divinationHandler.CreateDivination)
		authorized.GET("/history", divinationHandler.GetUserDivinations)
	}

	// 运势分析相关路由
	fortuneGroup := r.Group("/fortune").Use(middleware.Auth())
	{
		fortuneHandler := handlers.NewFortuneHandler()
		fortuneGroup.POST("/analyze", fortuneHandler.CalculateFortune)
		fortuneGroup.GET("/records", fortuneHandler.GetUserFortunes)
	}
}
