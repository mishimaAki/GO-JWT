package main

import (
	"GO-JWT/infrastructure/auth"
	"GO-JWT/interfaces/handler"
	"GO-JWT/internal/interface/middleware"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"
)

// 明日はgetEnvを実装するぞ！
func main() {
	r := gin.Default()

	limiter := middleware.NewIPRateLimiter(rate.Every(time.Minute/30), 30)
	r.Use(middleware.RateLimitMiddleware(limiter))

	jwtAuth := auth.NewJWTAuth("your-secret-key")
	authHandler := handler.NewAuthHandler(jwtAuth)

	r.POST("/login", authHandler.Login)

	r.Run(":8080")
}
